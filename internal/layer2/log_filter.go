package layer2

import (
	"context"
	"errors"
	"reflect"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/task"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (l *Layer2Listener) processLogs(vLog types.Log) {
	method, ok := l.addressBind[vLog.Address]
	if ok {
		err := method(vLog)
		if err != nil {
			log.Errorf("call %s processing log: %v", reflect.TypeOf(method).Name(), err)
		}
	}
}

func (l *Layer2Listener) processVotingLog(vLog types.Log) error {
	// save current submitter
	var (
		submitterChosen db.SubmitterChosen
		submitter       string
	)

	eventName := ""

	switch vLog.Topics[0] {
	case contracts.SubmitterChosenTopic:
		eventName = "SubmitterChosen"
		submitterChosenEvent := contracts.VotingManagerContractSubmitterChosen{}
		contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, eventName, vLog)
		submitter = submitterChosenEvent.NewSubmitter.Hex()
	case contracts.SubmitterRotationRequestedTopic:
		eventName = "SubmitterRotationRequested"
		submitterChosenEvent := contracts.VotingManagerContractSubmitterRotationRequested{}
		contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, eventName, vLog)
		submitter = submitterChosenEvent.CurrentSubmitter.Hex()
	}

	submitterChosen.BlockNumber = vLog.BlockNumber
	submitterChosen.Submitter = submitter
	submitterChosen.LogIndex = l.LogIndex(eventName, vLog)

	t := &task.SubmitterChosenPair{
		Old: db.SubmitterChosen{
			BlockNumber: l.state.TssState.BlockNumber,
			Submitter:   l.state.TssState.CurrentSubmitter.Hex(),
		},
		New: submitterChosen,
	}

	result := l.db.GetRelayerDB().Create(&submitterChosen)
	if result.RowsAffected > 0 {
		l.state.TssState.CurrentSubmitter = common.HexToAddress(submitter)
		l.state.TssState.BlockNumber = vLog.BlockNumber
		l.postTask(t)
	}

	return result.Error
}

func (l *Layer2Listener) processTaskLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.TaskSubmittedTopic:
		taskSubmitted := contracts.TaskManagerContractTaskSubmitted{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &taskSubmitted, "TaskSubmitted", vLog)
		task := &db.Task{
			TaskId:      uint32(taskSubmitted.TaskId.Uint64()),
			Context:     taskSubmitted.Context,
			Submitter:   taskSubmitted.Submitter.Hex(),
			BlockHeight: vLog.BlockNumber,
			LogIndex:    l.LogIndex("TaskSubmitted", vLog),
		}

		result := l.db.GetRelayerDB().Create(task)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected > 0 {
			l.postTask(task)
		}

	case contracts.TaskCompletedTopic:
		taskCompleted := contracts.TaskManagerContractTaskCompleted{}
		contracts.UnpackEventLog(contracts.TaskManagerContractMetaData, &taskCompleted, "TaskCompleted", vLog)

		err := l.db.GetRelayerDB().Transaction(func(tx *gorm.DB) error {
			taskErr := tx.
				Model(&db.Task{}).
				Where("task_id = ?", taskCompleted.TaskId.Uint64()).
				Update("status", db.Completed).Error
			err := tx.Save(l.LogIndex("TaskCompleted", vLog)).Error

			return errors.Join(taskErr, err)
		})
		if err != nil {
			return err
		}

		if l.state != nil && l.state.TssState.CurrentTask != nil {
			if uint32(taskCompleted.TaskId.Uint64()) >= l.state.TssState.CurrentTask.TaskId {
				l.state.TssState.CurrentTask = nil
			}
		}
	}

	return nil
}

func (l *Layer2Listener) LogIndex(eventName string, vlog types.Log) db.LogIndex {
	chainID, err := l.ChainID(context.Background())
	utils.Assert(err)

	if eventName == "" {
		eventName = vlog.Topics[0].String()
	}

	return db.LogIndex{
		ContractAddress: vlog.Address,
		EventName:       eventName,
		Log:             &vlog,
		TxHash:          vlog.TxHash,
		ChainId:         chainID.Uint64(),
		BlockNumber:     vlog.BlockNumber,
		LogIndex:        uint64(vlog.Index),
	}
}

func (l *Layer2Listener) processAccountLog(vLog types.Log) error {
	if vLog.Topics[0] == contracts.AddressRegisteredTopic {
		addressRegistered := contracts.AccountManagerContractAddressRegistered{}
		contracts.UnpackEventLog(contracts.AccountManagerContractMetaData, &addressRegistered, "AddressRegistered", vLog)
		account := db.Account{
			User:     addressRegistered.User.Hex(),
			Account:  addressRegistered.Account.Uint64(),
			ChainId:  addressRegistered.ChainId,
			Index:    addressRegistered.Index.Uint64(),
			Address:  addressRegistered.NewAddress.Hex(),
			LogIndex: l.LogIndex("AddressRegistered", vLog),
		}

		return l.db.GetRelayerDB().Create(&account).Error
	}

	return nil
}

func (l *Layer2Listener) processParticipantLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.ParticipantAddedTopic:
		eventParticipantAdded := contracts.ParticipantManagerContractParticipantAdded{}
		contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &eventParticipantAdded, "ParticipantAdded", vLog)
		newParticipant := eventParticipantAdded.Participant
		// save locked relayer member from db
		participant := db.Participant{
			Address: newParticipant.String(),
		}

		err := l.db.
			GetRelayerDB().Transaction(func(tx *gorm.DB) error {
			err1 := tx.FirstOrCreate(&participant, "address = ?", participant.Address).Error
			err2 := tx.Save(l.LogIndex("ParticipantAdded", vLog)).Error

			return errors.Join(err1, err2)
		})
		if err != nil {
			return err
		}

		log.Infof("Participant added: %s", newParticipant)

		if !slices.Contains(l.state.TssState.Participants, newParticipant) {
			pair := &task.ParticipantPair{Old: l.state.TssState.Participants}
			l.state.TssState.Participants = append(l.state.TssState.Participants, newParticipant)
			pair.New = l.state.TssState.Participants
			l.postTask(pair)
		}
	case contracts.ParticipantRemovedTopic:
		participantRemovedEvent := contracts.ParticipantManagerContractParticipantRemoved{}
		contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &participantRemovedEvent, "ParticipantRemoved", vLog)
		removedParticipant := participantRemovedEvent.Participant.Hex()

		err := l.db.
			GetRelayerDB().Transaction(func(tx *gorm.DB) error {
			removedErr := tx.Where("address = ?", removedParticipant).
				Delete(&db.Participant{}).
				Error
			vlogErr := tx.Save(l.LogIndex("ParticipantRemoved", vLog)).Error

			return errors.Join(removedErr, vlogErr)
		})
		if err != nil {
			return err
		}

		log.Infof("Participant removed: %s", removedParticipant)

		index := slices.Index(l.state.TssState.Participants, common.HexToAddress(removedParticipant))
		if index >= 0 {
			pair := &task.ParticipantPair{Old: l.state.TssState.Participants}
			l.state.TssState.Participants = slices.Delete(l.state.TssState.Participants, index, index)
			pair.New = l.state.TssState.Participants
			l.postTask(pair)
		}
	}

	return nil
}

func (l *Layer2Listener) processDepositLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.DepositRecordedTopic:
		depositRecorded := contracts.DepositManagerContractDepositRecorded{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &depositRecorded, "DepositRecorded", vLog)
		depositRecord := db.DepositRecord{
			TargetAddress: depositRecorded.TargetAddress.Hex(),
			Amount:        depositRecorded.Amount.Uint64(),
			ChainId:       depositRecorded.ChainId.Uint64(),
			TxInfo:        depositRecorded.TxInfo,
			ExtraInfo:     depositRecorded.ExtraInfo,
			LogIndex:      l.LogIndex("DepositRecorded", vLog),
		}

		return l.db.GetRelayerDB().Save(&depositRecord).Error
	case contracts.WithdrawalRecordedTopic:
		withdrawalRecorded := contracts.DepositManagerContractWithdrawalRecorded{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &withdrawalRecorded, "WithdrawalRecorded", vLog)
		withdrawalRecord := db.WithdrawalRecord{
			TargetAddress: withdrawalRecorded.TargetAddress.Hex(),
			Amount:        withdrawalRecorded.Amount.Uint64(),
			ChainId:       withdrawalRecorded.ChainId.Uint64(),
			TxInfo:        withdrawalRecorded.TxInfo,
			ExtraInfo:     withdrawalRecorded.ExtraInfo,
			LogIndex:      l.LogIndex("WithdrawalRecorded", vLog),
		}

		return l.db.GetRelayerDB().Save(&withdrawalRecord).Error
	}

	return nil
}
