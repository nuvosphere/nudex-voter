package layer2

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/abis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"slices"
	"time"
)

func (lis *Layer2Listener) processLogs(vLog types.Log) {
	if len(vLog.Topics) == 0 {
		log.Debug("No topics found in the log")
		return
	}
	var err error
	switch vLog.Address {
	case abis.VotingAddress:
		err = lis.processVotingLog(vLog)
	case abis.AccountAddress:
		err = lis.processAccountLog(vLog)
	case abis.OperationsAddress:
		err = lis.processOperationsLog(vLog)
	case abis.ParticipantAddress:
		err = lis.processParticipantLog(vLog)
	case abis.DepositAddress:
		err = lis.processDepositLog(vLog)
	}
	if err != nil {
		log.Errorf("Error processing log: %v", err)
	}
}

func (lis *Layer2Listener) processVotingLog(vLog types.Log) error {
	submitterChosenEvent, err := lis.contractVotingManager.ParseSubmitterChosen(vLog)
	if err == nil {
		// save current submitter
		var submitterChosen db.SubmitterChosen
		result := lis.db.GetRelayerDB().Where("block_number = ? AND submitter = ?",
			vLog.BlockNumber, submitterChosenEvent.NewSubmitter).First(&submitterChosen)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				submitterChosen.BlockNumber = vLog.BlockNumber
				submitterChosen.Submitter = submitterChosenEvent.NewSubmitter.Hex()
				err = lis.db.GetRelayerDB().Create(&submitterChosen).Error
				if err != nil {
					lis.state.TssState.CurrentSubmitter = submitterChosenEvent.NewSubmitter.Hex()
					lis.state.TssState.BlockNumber = vLog.BlockNumber
				}
			} else {
				return err
			}
		}
		return nil
	}
	return nil
}

func (lis *Layer2Listener) processOperationsLog(vLog types.Log) error {
	taskSubmitted, err := lis.contractOperations.ParseTaskSubmitted(vLog)
	if err == nil {
		var existingTask db.Task
		result := lis.db.GetRelayerDB().Where("task_id = ?", taskSubmitted.TaskId.Uint64()).First(&existingTask)

		if result.Error == nil {
			return nil
		} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			task := db.Task{
				TaskId:      taskSubmitted.TaskId.Uint64(),
				Context:     taskSubmitted.Context,
				Submitter:   taskSubmitted.Submitter.Hex(),
				IsCompleted: false,
				BlockHeight: vLog.BlockNumber,
				CreatedAt:   time.Now(),
			}
			err = lis.db.GetRelayerDB().Create(&task).Error
			if err != nil {
				return err
			}
			return nil
		} else {
			return result.Error
		}
	}

	taskCompleted, err := lis.contractOperations.ParseTaskCompleted(vLog)
	if err == nil {
		var existingTask db.Task
		result := lis.db.GetRelayerDB().Where("task_id = ?", taskCompleted.TaskId.Uint64()).First(&existingTask)
		if result.Error == nil {
			existingTask.IsCompleted = true
			existingTask.CompletedAt = time.Unix(taskCompleted.CompletedAt.Int64(), 0)
			err := lis.db.GetRelayerDB().Save(&existingTask).Error

			if lis.state != nil && &lis.state.TssState != nil && lis.state.TssState.CurrentTask != nil {
				if taskCompleted.TaskId.Uint64() >= lis.state.TssState.CurrentTask.TaskId {
					lis.state.TssState.CurrentTask = nil
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

func (lis *Layer2Listener) processAccountLog(vLog types.Log) error {
	addressRegistered, err := lis.contractAccountManager.ParseAddressRegistered(vLog)
	if err == nil {
		account := db.Account{
			User:    addressRegistered.User.Hex(),
			Account: addressRegistered.Account.Uint64(),
			ChainId: addressRegistered.ChainId,
			Index:   addressRegistered.Index.Uint64(),
			Address: addressRegistered.NewAddress.Hex(),
		}
		err = lis.db.GetRelayerDB().Create(&account).Error
		if err != nil {
			log.Fatalf("Error adding address: %v", err)
		} else {
			log.Infof("Address %s registered for user: %s, chain: %d", account.Address, account.User, account.ChainId)
		}
	}
	return nil
}

func (lis *Layer2Listener) processParticipantLog(vLog types.Log) error {
	eventParticipantAdded, err := lis.contractParticipantManager.ParseParticipantAdded(vLog)
	if err == nil {
		newParticipant := eventParticipantAdded.Participant.Hex()
		// save locked relayer member from db
		participant := db.Participant{
			Address: newParticipant,
		}
		err = lis.db.GetRelayerDB().FirstOrCreate(&participant, "address = ?", participant.Address).Error
		if err != nil {
			log.Fatalf("Error adding Participant: %v", err)
		} else {
			log.Infof("Participant added: %s", newParticipant)
			if !slices.Contains(lis.state.TssState.Participants, newParticipant) {
				lis.state.TssState.Participants = append(lis.state.TssState.Participants, newParticipant)
			}
		}
		return nil
	}

	participantRemovedEvent, err := lis.contractParticipantManager.ParseParticipantRemoved(vLog)
	if err == nil {
		removedParticipant := participantRemovedEvent.Participant.Hex()
		err = lis.db.GetRelayerDB().Where("address = ?",
			removedParticipant).Delete(&db.Participant{}).Error
		if err != nil {
			log.Fatalf("Error removing Participant: %v", err)
		} else {
			log.Infof("Participant removed: %s", removedParticipant)
			index := slices.Index(lis.state.TssState.Participants, removedParticipant)
			if index != -1 {
				lis.state.TssState.Participants = append(lis.state.TssState.Participants[:index], lis.state.TssState.Participants[index+1:]...)
			}
		}
		return nil
	}
	return nil
}

func (lis *Layer2Listener) processDepositLog(vLog types.Log) error {
	depositRecorded, err := lis.contractDepositManager.ParseDepositRecorded(vLog)
	if err == nil {
		depositRecord := db.DepositRecord{
			TargetAddress: depositRecorded.TargetAddress.Hex(),
			Amount:        depositRecorded.Amount.Uint64(),
			ChainId:       depositRecorded.ChainId.Uint64(),
			TxInfo:        depositRecorded.TxInfo,
			ExtraInfo:     depositRecorded.ExtraInfo,
		}
		err = lis.db.GetRelayerDB().FirstOrCreate(&depositRecord, "target_address = ? and amount = ? and chain_id = ? and tx_info = ?",
			depositRecorded.TargetAddress.Hex(), depositRecorded.Amount.Uint64(), depositRecorded.ChainId.Uint64(), depositRecorded.TxInfo,
		).Error
		if err != nil {
			return err
		}
		return nil
	}

	withdrawalRecorded, err := lis.contractDepositManager.ParseWithdrawalRecorded(vLog)
	if err == nil {
		withdrawalRecord := db.WithdrawalRecord{
			TargetAddress: withdrawalRecorded.TargetAddress.Hex(),
			Amount:        withdrawalRecorded.Amount.Uint64(),
			ChainId:       withdrawalRecorded.ChainId.Uint64(),
			TxInfo:        withdrawalRecorded.TxInfo,
			ExtraInfo:     withdrawalRecorded.ExtraInfo,
		}
		err = lis.db.GetRelayerDB().FirstOrCreate(&withdrawalRecord, "target_address = ? and amount = ? and chain_id = ? and tx_info = ?",
			withdrawalRecorded.TargetAddress.Hex(), withdrawalRecorded.Amount.Uint64(), withdrawalRecorded.ChainId.Uint64(), withdrawalRecorded.TxInfo,
		).Error
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
