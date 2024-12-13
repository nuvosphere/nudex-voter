package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	SubmitterRotationRequestedTopic = crypto.Keccak256Hash([]byte(`SubmitterRotationRequested(address,address)`))                                                              // VotingManagerUpgradeable
	SubmitterChosenTopic            = crypto.Keccak256Hash([]byte(`SubmitterChosen(address)`))                                                                                 // VotingManagerUpgradeable
	TaskSubmittedTopic              = crypto.Keccak256Hash([]byte(`TaskSubmitted(uint64,bytes,address)`))                                                                      // TaskManager
	TaskUpdatedTopic                = crypto.Keccak256Hash([]byte(`TaskUpdated(uint64,address,uint256,bytes)`))                                                                // TaskManager
	AddressRegisteredTopic          = crypto.Keccak256Hash([]byte(`AddressRegistered(address,uint256,uint8,uint256,address)`))                                                 // AccountManagerContract
	ParticipantAddedTopic           = crypto.Keccak256Hash([]byte(`ParticipantAdded(address)`))                                                                                // IParticipantManager
	ParticipantRemovedTopic         = crypto.Keccak256Hash([]byte(`ParticipantRemoved(address)`))                                                                              // IParticipantManager
	DepositRecordedTopic            = crypto.Keccak256Hash([]byte(`DepositRecorded(address,uint256,uint256,bytes,bytes)`))                                                     // DepositManagerContract
	WithdrawalRecordedTopic         = crypto.Keccak256Hash([]byte(`WithdrawalRecorded(address,uint256,uint256,bytes,bytes)`))                                                  // DepositManagerContract
	NIP20TokenEventBurnbTopic       = crypto.Keccak256Hash([]byte(`NIP20TokenEvent_burnb(address,bytes32,uint256)`))                                                           // InscriptionContract
	NIP20TokenEventMintbTopic       = crypto.Keccak256Hash([]byte(`NIP20TokenEvent_mintb(address,bytes32,uint256)`))                                                           // InscriptionContract
	WalletCreationRequestTopic      = crypto.Keccak256Hash([]byte(`WalletCreationRequest(uint8,uint8,uint32,uint8,uint8)`))                                                    // TopicPayloadContract
	DepositRequestTopic             = crypto.Keccak256Hash([]byte(`DepositRequest(uint8,uint8,string,uint64,uint8,uint32,uint64,string,string,string,uint8,uint8)`))           // TopicPayloadContract
	WithdrawalRequestTopic          = crypto.Keccak256Hash([]byte(`WithdrawalRequest(uint8,uint8,string,uint64,uint8,uint32,uint64,string,string,string,uint8,uint8,uint64)`)) // TopicPayloadContract
	WalletCreationResultTopic       = crypto.Keccak256Hash([]byte(`WalletCreationResult(uint8,bool,uint8,string)`))                                                            // TopicPayloadContract
	DepositResultTopic              = crypto.Keccak256Hash([]byte(`DepositResult(uint8,bool,uint8)`))                                                                          // TopicPayloadContract
	WithdrawalResultTopic           = crypto.Keccak256Hash([]byte(`WithdrawalResult(uint8,bool,uint8)`))                                                                       // TopicPayloadContract
)

var Topics = [][]common.Hash{
	{SubmitterRotationRequestedTopic},
	{SubmitterChosenTopic},
	{TaskSubmittedTopic},
	{TaskUpdatedTopic},
	{AddressRegisteredTopic},
	{ParticipantRemovedTopic},
	{ParticipantAddedTopic},
	{DepositRecordedTopic},
	{WithdrawalRecordedTopic},
}
