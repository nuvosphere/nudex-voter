package layer2

import (
	"errors"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/state"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (l *Layer2Listener) processLogs(vLog types.Log) {
	method, ok := l.addressBind[vLog.Address]
	if ok {
		err := method(vLog)
		if err != nil {
			log.Errorf("Error processing log: %v", err)
		}
	}
}

func (l *Layer2Listener) processVotingLog(vLog types.Log) error {
	// save current submitter
	var (
		submitterChosen db.SubmitterChosen
		submitter       string
		err             error
	)

	switch vLog.Topics[0] {
	case SubmitterChosenTopic:
		submitterChosenEvent := contracts.VotingManagerContractSubmitterChosen{}
		err = contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, "SubmitterChosen", vLog)
		submitter = submitterChosenEvent.NewSubmitter.Hex()
	case SubmitterRotationRequestedTopic:
		submitterChosenEvent := contracts.VotingManagerContractSubmitterRotationRequested{}
		err = contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, "SubmitterRotationRequested", vLog)
		submitter = submitterChosenEvent.CurrentSubmitter.Hex()
	}

	if err != nil {
		return err
	}

	result := l.db.GetRelayerDB().
		Where("block_number = ? AND submitter = ?", vLog.BlockNumber, submitter).
		First(&submitterChosen)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			submitterChosen.BlockNumber = vLog.BlockNumber
			submitterChosen.Submitter = submitter

			err = l.db.GetRelayerDB().Create(&submitterChosen).Error
			if err != nil {
				l.state.TssState.CurrentSubmitter = common.HexToAddress(submitter)
				l.state.TssState.BlockNumber = vLog.BlockNumber
			}
		} else {
			return err
		}
	}

	return nil
}

func (l *Layer2Listener) processOperationsLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case TaskSubmittedTopic:
		taskSubmitted := contracts.TaskManagerContractTaskSubmitted{}

		err := contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &taskSubmitted, "TaskSubmitted", vLog)
		if err != nil {
			return err
		}

		var existingTask db.Task
		result := l.db.GetRelayerDB().Where("task_id = ?", taskSubmitted.TaskId.Uint64()).First(&existingTask)

		if result.Error == nil {
			return nil
		} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			task := db.Task{
				TaskId:      uint32(taskSubmitted.TaskId.Uint64()),
				Context:     taskSubmitted.Context,
				Submitter:   taskSubmitted.Submitter.Hex(),
				IsCompleted: false,
				BlockHeight: vLog.BlockNumber,
				CreatedAt:   time.Now(),
			}

			err = l.db.GetRelayerDB().Create(&task).Error
			if err != nil {
				return err
			}

			return nil
		} else {
			return result.Error
		}

	case TaskCompletedTopic:
		taskCompleted := contracts.TaskManagerContractTaskCompleted{}

		err := contracts.UnpackEventLog(contracts.TaskManagerContractMetaData, &taskCompleted, "TaskCompleted", vLog)
		if err != nil {
			return err
		}

		var existingTask db.Task

		result := l.db.GetRelayerDB().Where("task_id = ?", taskCompleted.TaskId.Uint64()).First(&existingTask)
		if result.Error == nil {
			existingTask.IsCompleted = true
			existingTask.CompletedAt = time.Unix(taskCompleted.CompletedAt.Int64(), 0)
			err := l.db.GetRelayerDB().Save(&existingTask).Error

			if l.state != nil && l.state.TssState.CurrentTask != nil {
				if uint32(taskCompleted.TaskId.Uint64()) >= l.state.TssState.CurrentTask.TaskId {
					l.state.TssState.CurrentTask = nil
				}
			}

			if err != nil {
				return err
			}
		} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Fatalf("Task %d not found for event TaskCompleted: %v", taskCompleted.TaskId.Uint64(), taskCompleted)
		}
	}

	return nil
}

func (l *Layer2Listener) processAccountLog(vLog types.Log) error {
	if vLog.Topics[0] == AddressRegisteredTopic {
		addressRegistered := contracts.AccountManagerContractAddressRegistered{}

		err := contracts.UnpackEventLog(contracts.AccountManagerContractMetaData, &addressRegistered, "AddressRegistered", vLog)
		if err != nil {
			return err
		}

		account := db.Account{
			User:    addressRegistered.User.Hex(),
			Account: addressRegistered.Account.Uint64(),
			ChainId: addressRegistered.ChainId,
			Index:   addressRegistered.Index.Uint64(),
			Address: addressRegistered.NewAddress.Hex(),
		}

		err = l.db.GetRelayerDB().Create(&account).Error
		if err != nil {
			log.Fatalf("Error adding address: %v", err)
		} else {
			log.Infof("localAddress %s registered for user: %s, chain: %d", account.Address, account.User, account.ChainId)
		}
	}

	return nil
}

func (l *Layer2Listener) processParticipantLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case ParticipantAddedTopic:
		eventParticipantAdded := contracts.ParticipantManagerContractParticipantAdded{}

		err := contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &eventParticipantAdded, "ParticipantAdded", vLog)
		if err != nil {
			return err
		}

		newParticipant := eventParticipantAdded.Participant
		// save locked relayer member from db
		participant := db.Participant{Address: newParticipant.String()}

		err = l.db.
			GetRelayerDB().
			FirstOrCreate(&participant, "address = ?", participant.Address).
			Error
		if err != nil {
			return err
		}

		log.Infof("Participant added: %s", newParticipant)

		if !slices.Contains(l.state.TssState.Participants, newParticipant) {
			l.state.TssState.Participants = append(l.state.TssState.Participants, newParticipant)
			l.state.EventBus.Publish(state.EventParticipantAddedOrRemoved{}, l.state.TssState.Participants)
		}
	case ParticipantRemovedTopic:
		participantRemovedEvent := contracts.ParticipantManagerContractParticipantRemoved{}

		err := contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &participantRemovedEvent, "ParticipantRemoved", vLog)
		if err != nil {
			return err
		}

		removedParticipant := participantRemovedEvent.Participant.Hex()

		err = l.db.
			GetRelayerDB().
			Where("address = ?", removedParticipant).
			Delete(&db.Participant{}).
			Error
		if err != nil {
			return err
		}

		log.Infof("Participant removed: %s", removedParticipant)

		index := slices.Index(l.state.TssState.Participants, common.HexToAddress(removedParticipant))
		if index >= 0 {
			l.state.TssState.Participants = slices.Delete(l.state.TssState.Participants, index, index)
			l.state.EventBus.Publish(state.EventParticipantAddedOrRemoved{}, l.state.TssState.Participants)
		}
	}

	return nil
}

func (l *Layer2Listener) processDepositLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case DepositRecordedTopic:
		depositRecorded := contracts.DepositManagerContractDepositRecorded{}

		err := contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &depositRecorded, "DepositRecorded", vLog)
		if err != nil {
			return err
		}

		depositRecord := db.DepositRecord{
			TargetAddress: depositRecorded.TargetAddress.Hex(),
			Amount:        depositRecorded.Amount.Uint64(),
			ChainId:       depositRecorded.ChainId.Uint64(),
			TxInfo:        depositRecorded.TxInfo,
			ExtraInfo:     depositRecorded.ExtraInfo,
		}

		return l.db.GetRelayerDB().FirstOrCreate(&depositRecord, "target_address = ? and amount = ? and chain_id = ? and tx_info = ?",
			depositRecorded.TargetAddress.Hex(), depositRecorded.Amount.Uint64(), depositRecorded.ChainId.Uint64(), depositRecorded.TxInfo,
		).Error
	case WithdrawalRecordedTopic:
		withdrawalRecorded := contracts.DepositManagerContractWithdrawalRecorded{}

		err := contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &withdrawalRecorded, "WithdrawalRecorded", vLog)
		if err != nil {
			return err
		}

		withdrawalRecord := db.WithdrawalRecord{
			TargetAddress: withdrawalRecorded.TargetAddress.Hex(),
			Amount:        withdrawalRecorded.Amount.Uint64(),
			ChainId:       withdrawalRecorded.ChainId.Uint64(),
			TxInfo:        withdrawalRecorded.TxInfo,
			ExtraInfo:     withdrawalRecorded.ExtraInfo,
		}

		return l.db.GetRelayerDB().FirstOrCreate(&withdrawalRecord, "target_address = ? and amount = ? and chain_id = ? and tx_info = ?",
			withdrawalRecorded.TargetAddress.Hex(), withdrawalRecorded.Amount.Uint64(), withdrawalRecorded.ChainId.Uint64(), withdrawalRecorded.TxInfo,
		).Error
	}

	return nil
}
