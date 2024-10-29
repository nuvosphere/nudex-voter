package layer2

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
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
	}
	if err != nil {
		log.Errorf("Error processing log: %v", err)
	}
}

func (lis *Layer2Listener) processVotingLog(vLog types.Log) error {
	eventSubmitterRotation, err := lis.contractVotingManager.ParseSubmitterRotationRequested(vLog)
	if err == nil {
		// save current rotation
		var submitterRotation db.SubmitterRotation
		result := lis.db.GetRelayerDB().First(&submitterRotation)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				submitterRotation.BlockNumber = vLog.BlockNumber
				submitterRotation.CurrentSubmitter = eventSubmitterRotation.CurrentSubmitter.Hex()
				err = lis.db.GetRelayerDB().Create(&submitterRotation).Error
				if err != nil {
					lis.state.TssState.CurrentSubmitter = eventSubmitterRotation.CurrentSubmitter.Hex()
					lis.state.TssState.BlockNumber = vLog.BlockNumber
				}
			} else {
				log.Fatalf("Error query db: %v", result.Error)
			}
		} else {
			submitterRotation.BlockNumber = vLog.BlockNumber
			submitterRotation.CurrentSubmitter = eventSubmitterRotation.CurrentSubmitter.Hex()
			err = lis.db.GetRelayerDB().Save(&submitterRotation).Error
			if err != nil {
				lis.state.TssState.CurrentSubmitter = eventSubmitterRotation.CurrentSubmitter.Hex()
				lis.state.TssState.BlockNumber = vLog.BlockNumber
			}
		}

		if err != nil {
			log.Fatalf("Error saving SubmitterRotation: %v", err)
		}

		selfAddress := crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey)
		if eventSubmitterRotation.CurrentSubmitter == selfAddress {
			// TODO start submitDeposit call
		} else {
			// TODO stop submitDeposit
		}
		return nil
	}

	eventParticipantAdded, err := lis.contractVotingManager.ParseParticipantAdded(vLog)
	if err == nil {
		newParticipant := eventParticipantAdded.NewParticipant.Hex()
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

	eventParticipantRemoved, err := lis.contractVotingManager.ParseParticipantRemoved(vLog)
	if err == nil {
		removedParticipant := eventParticipantRemoved.Participant.Hex()
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
