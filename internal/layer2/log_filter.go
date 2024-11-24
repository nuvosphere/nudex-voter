package layer2

import (
	"context"
	"errors"
	"reflect"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
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
		eventName = SubmitterChosen
		submitterChosenEvent := contracts.VotingManagerContractSubmitterChosen{}
		contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, eventName, vLog)
		submitter = submitterChosenEvent.NewSubmitter.Hex()
	case contracts.SubmitterRotationRequestedTopic:
		eventName = SubmitterRotationRequested
		submitterChosenEvent := contracts.VotingManagerContractSubmitterRotationRequested{}
		contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, eventName, vLog)
		submitter = submitterChosenEvent.CurrentSubmitter.Hex()
	}

	submitterChosen.Submitter = submitter
	submitterChosen.BlockNumber = vLog.BlockNumber
	submitterChosen.LogIndex = l.LogIndex(eventName, vLog)

	result := l.db.GetRelayerDB().Create(&submitterChosen)
	if result.RowsAffected > 0 {
		l.postTask(submitterChosen)
	}

	return result.Error
}

func (l *Layer2Listener) processTaskLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.TaskSubmittedTopic:
		taskSubmitted := contracts.TaskManagerContractTaskSubmitted{}
		contracts.UnpackEventLog(contracts.TaskManagerContractMetaData, &taskSubmitted, TaskSubmitted, vLog)
		actualTask := db.DecodeTask(taskSubmitted.TaskId, taskSubmitted.Context)
		task := db.Task{
			TaskId:    actualTask.TaskID(),
			TaskType:  actualTask.Type(),
			Context:   taskSubmitted.Context,
			Submitter: taskSubmitted.Submitter.Hex(),
			LogIndex:  l.LogIndex(TaskSubmitted, vLog),
		}
		actualTask.SetBaseTask(task)

		result := l.db.GetRelayerDB().Create(actualTask)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected > 0 {
			l.postTask(actualTask)
		}

	case contracts.TaskUpdatedTopic:
		taskUpdated := contracts.TaskManagerContractTaskUpdated{}
		contracts.UnpackEventLog(contracts.TaskManagerContractMetaData, &taskUpdated, TaskUpdated, vLog)

		var taskUpdatedEvent *db.TaskUpdatedEvent

		err := l.db.GetRelayerDB().Transaction(func(tx *gorm.DB) error {
			taskErr := tx.
				Model(&db.Task{}).
				Where("task_id = ?", taskUpdated.TaskId).
				Update("status", db.Completed).Error

			taskUpdatedEvent = &db.TaskUpdatedEvent{
				TaskId:     taskUpdated.TaskId,
				Submitter:  taskUpdated.Submitter.Hex(),
				UpdateTime: taskUpdated.UpdateTime.Int64(),
				Result:     taskUpdated.Result,
				LogIndex:   l.LogIndex(TaskUpdated, vLog),
			}
			err := tx.Save(taskUpdatedEvent).Error

			return errors.Join(taskErr, err)
		})
		if err != nil {
			return err
		}

		if taskUpdatedEvent != nil {
			l.postTask(taskUpdatedEvent)
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
		contracts.UnpackEventLog(contracts.AccountManagerContractMetaData, &addressRegistered, AddressRegistered, vLog)
		account := db.Account{
			Account:  addressRegistered.Account.Uint64(),
			ChainId:  addressRegistered.ChainId,
			Index:    addressRegistered.Index.Uint64(),
			Address:  addressRegistered.NewAddress.Hex(),
			LogIndex: l.LogIndex(AddressRegistered, vLog),
		}

		return l.db.GetRelayerDB().Create(&account).Error
	}

	return nil
}

func (l *Layer2Listener) processParticipantLog(vLog types.Log) error {
	var (
		participantEvent *db.ParticipantEvent
		err              error
	)

	switch vLog.Topics[0] {
	case contracts.ParticipantAddedTopic:
		eventParticipantAdded := contracts.ParticipantManagerContractParticipantAdded{}
		contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &eventParticipantAdded, ParticipantAdded, vLog)
		newParticipant := eventParticipantAdded.Participant
		// save locked relayer member from db
		participant := db.Participant{Address: newParticipant.String()}
		err = l.db.
			GetRelayerDB().Transaction(func(tx *gorm.DB) error {
			err1 := tx.FirstOrCreate(&participant, "address = ?", participant.Address).Error
			participantEvent = &db.ParticipantEvent{
				EventName:   ParticipantAdded,
				Address:     participant.Address,
				BlockNumber: vLog.BlockNumber,
				LogIndex:    l.LogIndex(ParticipantAdded, vLog),
			}
			err2 := tx.Save(participantEvent).Error

			return errors.Join(err1, err2)
		})
	case contracts.ParticipantRemovedTopic:
		participantRemovedEvent := contracts.ParticipantManagerContractParticipantRemoved{}
		contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &participantRemovedEvent, ParticipantRemoved, vLog)
		removedParticipant := participantRemovedEvent.Participant.Hex()

		err = l.db.
			GetRelayerDB().Transaction(func(tx *gorm.DB) error {
			removedErr := tx.Where("address = ?", removedParticipant).
				Delete(&db.Participant{}).
				Error
			participantEvent := &db.ParticipantEvent{
				EventName:   ParticipantRemoved,
				Address:     removedParticipant,
				BlockNumber: vLog.BlockNumber,
				LogIndex:    l.LogIndex(ParticipantRemoved, vLog),
			}
			vlogErr := tx.Save(participantEvent).Error

			return errors.Join(removedErr, vlogErr)
		})
	}

	if err != nil {
		return err
	}

	if participantEvent != nil {
		l.postTask(participantEvent)
		log.Infof("Participant %s: %s", participantEvent.EventName, participantEvent.Address)
	}

	return nil
}

func (l *Layer2Listener) processDepositLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.DepositRecordedTopic:
		depositRecorded := contracts.DepositManagerContractDepositRecorded{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &depositRecorded, DepositRecorded, vLog)
		depositRecord := db.DepositRecord{
			TargetAddress: depositRecorded.TargetAddress.Hex(),
			Amount:        depositRecorded.Amount.Uint64(),
			ChainId:       depositRecorded.ChainId.Uint64(),
			TxInfo:        depositRecorded.TxInfo,
			ExtraInfo:     depositRecorded.ExtraInfo,
			LogIndex:      l.LogIndex(DepositRecorded, vLog),
		}

		return l.db.GetRelayerDB().Save(&depositRecord).Error
	case contracts.WithdrawalRecordedTopic:
		withdrawalRecorded := contracts.DepositManagerContractWithdrawalRecorded{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &withdrawalRecorded, WithdrawalRecorded, vLog)
		withdrawalRecord := db.WithdrawalRecord{
			TargetAddress: withdrawalRecorded.TargetAddress.Hex(),
			Amount:        withdrawalRecorded.Amount.Uint64(),
			ChainId:       withdrawalRecorded.ChainId.Uint64(),
			TxInfo:        withdrawalRecorded.TxInfo,
			ExtraInfo:     withdrawalRecorded.ExtraInfo,
			LogIndex:      l.LogIndex(WithdrawalRecorded, vLog),
		}

		return l.db.GetRelayerDB().Save(&withdrawalRecord).Error
	}

	return nil
}
