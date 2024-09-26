package layer2

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func processVotingLog(lis *Layer2Listener, vLog types.Log, parsedABI abi.ABI) {
	// TODO  DepositInfoSubmitted event parse and update deposit tx in db
	switch vLog.Topics[0].Hex() {
	case parsedABI.Events["SubmitterRotationRequested"].ID.Hex():
		event := struct {
			Requester        common.Address
			CurrentSubmitter common.Address
		}{}

		err := parsedABI.UnpackIntoInterface(&event, "SubmitterRotationRequested", vLog.Data)
		if err != nil {
			log.Errorf("Error unpacking SubmitterRotationRequested event: %v", err)
		}

		// save current rotation
		var submitterRotation db.SubmitterRotation
		result := lis.db.GetRelayerDB().First(&submitterRotation)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				submitterRotation.BlockNumber = vLog.BlockNumber
				submitterRotation.CurrentSubmitter = event.CurrentSubmitter.Hex()
				err = lis.db.GetRelayerDB().Create(&submitterRotation).Error
			} else {
				log.Fatalf("Error query db: %v", result.Error)
			}
		} else {
			submitterRotation.BlockNumber = vLog.BlockNumber
			submitterRotation.CurrentSubmitter = event.CurrentSubmitter.Hex()
			err = lis.db.GetRelayerDB().Save(&submitterRotation).Error
		}

		if err != nil {
			log.Fatalf("Error saving SubmitterRotation: %v", err)
		}

		selfAddress := crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey)
		if event.CurrentSubmitter == selfAddress {
			// TODO start submitDeposit call
		} else {
			// TODO stop submitDeposit
		}

	case parsedABI.Events["ParticipantAdded"].ID.Hex():
		event := struct {
			NewParticipant common.Address
		}{}

		err := parsedABI.UnpackIntoInterface(&event, "ParticipantAdded", vLog.Data)
		if err != nil {
			log.Errorf("Error unpacking ParticipantAdded event: %v", err)
			return
		}

		// save locked relayer member from db
		participant := db.Participant{
			Address: event.NewParticipant.Hex(),
		}
		err = lis.db.GetRelayerDB().FirstOrCreate(&participant, "address = ?", participant.Address).Error
		if err != nil {
			log.Fatalf("Error adding Participant: %v", err)
		} else {
			log.Infof("Participant added: %s", event.NewParticipant.Hex())
		}

	case parsedABI.Events["ParticipantRemoved"].ID.Hex():
		event := struct {
			Participant common.Address
		}{}

		err := parsedABI.UnpackIntoInterface(&event, "ParticipantRemoved", vLog.Data)
		if err != nil {
			log.Errorf("Error unpacking ParticipantRemoved event: %v", err)
			return
		}

		// remove relayer member from db
		err = lis.db.GetRelayerDB().Where("address = ?", event.Participant.Hex()).Delete(&db.Participant{}).Error
		if err != nil {
			log.Fatalf("Error removing Participant: %v", err)
		} else {
			log.Infof("Participant removed: %s", event.Participant.Hex())
		}
	}
}
