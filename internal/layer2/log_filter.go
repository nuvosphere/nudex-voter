package layer2

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/abis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
			if result.Error == gorm.ErrRecordNotFound {
				submitterRotation.BlockNumber = vLog.BlockNumber
				submitterRotation.CurrentSubmitter = eventSubmitterRotation.CurrentSubmitter.Hex()
				err = lis.db.GetRelayerDB().Create(&submitterRotation).Error
			} else {
				log.Fatalf("Error query db: %v", result.Error)
			}
		} else {
			submitterRotation.BlockNumber = vLog.BlockNumber
			submitterRotation.CurrentSubmitter = eventSubmitterRotation.CurrentSubmitter.Hex()
			err = lis.db.GetRelayerDB().Save(&submitterRotation).Error
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
		// save locked relayer member from db
		participant := db.Participant{
			Address: eventParticipantAdded.NewParticipant.Hex(),
		}
		err = lis.db.GetRelayerDB().FirstOrCreate(&participant, "address = ?", participant.Address).Error
		if err != nil {
			log.Fatalf("Error adding Participant: %v", err)
		} else {
			log.Infof("Participant added: %s", eventParticipantAdded.NewParticipant.Hex())
		}
		return nil
	}

	eventParticipantRemoved, err := lis.contractVotingManager.ParseParticipantRemoved(vLog)
	if err == nil {
		err = lis.db.GetRelayerDB().Where("address = ?",
			eventParticipantRemoved.Participant.Hex()).Delete(&db.Participant{}).Error
		if err != nil {
			log.Fatalf("Error removing Participant: %v", err)
		} else {
			log.Infof("Participant removed: %s", eventParticipantRemoved.Participant.Hex())
		}
		return nil
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
