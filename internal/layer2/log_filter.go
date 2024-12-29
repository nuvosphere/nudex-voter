package layer2

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	default:
		return errors.New("invalid topic")
	}

	submitterChosen.Submitter = submitter
	submitterChosen.BlockNumber = vLog.BlockNumber
	submitterChosen.LogIndex = l.LogIndex(eventName, vLog)

	result := l.db.GetL2InfoDB().Create(&submitterChosen)
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
		actualTask := DecodeTask(taskSubmitted.TaskId, taskSubmitted.Data)
		baseTask := db.Task{
			TaskId:    actualTask.TaskID(),
			TaskType:  actualTask.Type(),
			Context:   taskSubmitted.Data,
			Submitter: taskSubmitted.Submitter.Hex(),
			LogIndex:  l.LogIndex(TaskSubmitted, vLog),
		}
		actualTask.SetBaseTask(baseTask)

		result := l.db.GetL2InfoDB().Create(actualTask)
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

		err := l.db.GetL2InfoDB().Transaction(func(tx *gorm.DB) error {
			task := db.Task{}
			taskErr := tx.
				Model(&task).
				Clauses(clause.Returning{}).
				Where("task_id = ?", taskUpdated.TaskId).
				Update("status", taskUpdated.State).
				Error

			taskUpdatedEvent = &db.TaskUpdatedEvent{
				TaskId:     taskUpdated.TaskId,
				Submitter:  taskUpdated.Submitter.Hex(),
				UpdateTime: taskUpdated.UpdateTime.Int64(),
				State:      taskUpdated.State,
				Result:     taskUpdated.Result,
				LogIndex:   l.LogIndex(TaskUpdated, vLog),
			}
			err := tx.Save(taskUpdatedEvent).Error
			taskUpdatedEvent.Task = task

			return errors.Join(taskErr, err)
		})
		if err != nil {
			return err
		}

		if taskUpdatedEvent != nil {
			l.postTask(taskUpdatedEvent)
		}
	case contracts.NIP20TokenEventMintbTopic:
		mintb := contracts.InscriptionContractNIP20TokenEventMintb{}
		contracts.UnpackEventLog(contracts.InscriptionContractMetaData, &mintb, NIP20TokenMintbEvent, vLog)

		mintbEvent := &db.InscriptionMintb{
			Recipient: mintb.Recipient.Hex(),
			Ticker:    mintb.Ticker,
			Amount:    decimal.NewFromBigInt(mintb.Amount, 0),
			LogIndex:  l.LogIndex(NIP20TokenMintbEvent, vLog),
		}
		err := l.db.GetL2InfoDB().Create(mintbEvent)
		if err != nil {
			return err.Error
		}
	case contracts.NIP20TokenEventBurnbTopic:
		burnb := contracts.InscriptionContractNIP20TokenEventBurnb{}
		contracts.UnpackEventLog(contracts.InscriptionContractMetaData, &burnb, NIP20TokenBurnbEvent, vLog)

		burnbEvent := &db.InscriptionBurnb{
			From:     burnb.From.Hex(),
			Ticker:   burnb.Ticker,
			Amount:   decimal.NewFromBigInt(burnb.Amount, 0),
			LogIndex: l.LogIndex(NIP20TokenBurnbEvent, vLog),
		}
		err := l.db.GetL2InfoDB().Create(burnbEvent)
		if err != nil {
			return err.Error
		}
	default:
		return errors.New("invalid topic")
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
		ChainId:         vtypes.BigToByte32(chainID),
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
			Chain:    addressRegistered.Chain,
			Index:    uint32(addressRegistered.Index.Uint64()),
			Address:  addressRegistered.NewAddress,
			LogIndex: l.LogIndex(AddressRegistered, vLog),
		}

		return l.db.GetL2InfoDB().Create(&account).Error
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
			GetL2InfoDB().Transaction(func(tx *gorm.DB) error {
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
			GetL2InfoDB().Transaction(func(tx *gorm.DB) error {
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
	default:
		return errors.New("invalid topic")
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
		err := l.db.GetL2InfoDB().Transaction(func(tx *gorm.DB) error {
			depositRecord := db.DepositRecord{
				TargetAddress: strings.ToLower(depositRecorded.DepositAddress),
				Amount:        decimal.NewFromBigInt(depositRecorded.Amount, 0),
				ChainId:       depositRecorded.ChainId,
				LogIndex:      l.LogIndex(DepositRecorded, vLog),
			}
			err1 := tx.Save(&depositRecord).Error

			addressBalance := db.AddressBalance{
				Address: depositRecorded.DepositAddress,
				// Token:   ,
				Amount:  decimal.NewFromBigInt(depositRecorded.Amount, 0),
				ChainId: depositRecorded.ChainId,
			}

			err2 := tx.Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "address"}},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"amount": gorm.Expr("amount + ?", addressBalance.Amount),
				}),
			}).Create(&addressBalance).Error

			return errors.Join(err1, err2)
		})

		return err
	case contracts.WithdrawalRecordedTopic:
		withdrawalRecorded := contracts.DepositManagerContractWithdrawalRecorded{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &withdrawalRecorded, WithdrawalRecorded, vLog)

		err := l.db.GetL2InfoDB().Transaction(func(tx *gorm.DB) error {
			withdrawalRecord := db.WithdrawalRecord{
				DepositAddress: withdrawalRecorded.DepositAddress,
				Amount:         decimal.NewFromBigInt(withdrawalRecorded.Amount, 0),
				ChainId:        withdrawalRecorded.ChainId,
				LogIndex:       l.LogIndex(WithdrawalRecorded, vLog),
			}
			err1 := tx.Save(&withdrawalRecord).Error

			err2 := tx.Model(&db.AddressBalance{}).
				Where("address = ?", withdrawalRecorded.DepositAddress).
				Update("amount", gorm.Expr("amount - ?", decimal.NewFromBigInt(withdrawalRecorded.Amount, 0))).
				Error

			return errors.Join(err1, err2)
		})

		return err
	}

	return nil
}

func (l *Layer2Listener) processAssetLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.AssetListedTopic:
		event := contracts.AssetHandlerContractAssetListed{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, AssetListed, vLog)

		asset := &db.Asset{
			Ticker:            event.Ticker,
			Decimals:          event.AssetParam.Decimals,
			DepositEnabled:    event.AssetParam.DepositEnabled,
			WithdrawalEnabled: event.AssetParam.WithdrawalEnabled,
			MinDepositAmount:  event.AssetParam.MinDepositAmount.Uint64(),
			MinWithdrawAmount: event.AssetParam.MinWithdrawAmount.Uint64(),
			AssetAlias:        event.AssetParam.AssetAlias,
		}
		return l.db.GetL2SyncDB().Save(asset).Error
	case contracts.AssetUpdatedTopic:
		event := contracts.AssetHandlerContractAssetUpdated{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, AssetUpdated, vLog)

		var existingAsset db.Asset
		result := l.db.GetL2SyncDB().Where("ticker = ?", event.Ticker).First(&existingAsset)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				newAsset := &db.Asset{
					Ticker:            event.Ticker,
					Decimals:          event.AssetParam.Decimals,
					DepositEnabled:    event.AssetParam.DepositEnabled,
					WithdrawalEnabled: event.AssetParam.WithdrawalEnabled,
					MinDepositAmount:  event.AssetParam.MinDepositAmount.Uint64(),
					MinWithdrawAmount: event.AssetParam.MinWithdrawAmount.Uint64(),
					AssetAlias:        event.AssetParam.AssetAlias,
				}
				return l.db.GetL2SyncDB().Save(newAsset).Error
			}
			return fmt.Errorf("failed to query asset by ticker:%s, %w", event.Ticker, result.Error)
		}

		existingAsset.Decimals = event.AssetParam.Decimals
		existingAsset.DepositEnabled = event.AssetParam.DepositEnabled
		existingAsset.WithdrawalEnabled = event.AssetParam.WithdrawalEnabled
		existingAsset.MinDepositAmount = event.AssetParam.MinDepositAmount.Uint64()
		existingAsset.MinWithdrawAmount = event.AssetParam.MinWithdrawAmount.Uint64()
		existingAsset.AssetAlias = event.AssetParam.AssetAlias

		return l.db.GetL2SyncDB().Save(&existingAsset).Error
	case contracts.AssetDelistedTopic:
		event := contracts.AssetHandlerContractAssetDelisted{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, AssetDelisted, vLog)

		ticker := strings.Trim(string(event.Ticker[:]), "\x00")

		var existingAsset db.Asset
		result := l.db.GetL2SyncDB().Where("ticker = ?", ticker).First(&existingAsset)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return fmt.Errorf("failed to query asset by ticker:%s, %w", ticker, result.Error)
		}

		if err := l.db.GetL2SyncDB().Delete(&existingAsset).Error; err != nil {
			return fmt.Errorf("failed to delete asset with ticker: %s, %w", ticker, err)
		}

		return nil
	default:
		return errors.New("invalid topic")
	}
}
