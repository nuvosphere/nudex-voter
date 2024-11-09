package layer2

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	SubmitterRotationRequestedTopic = crypto.Keccak256Hash([]byte(`SubmitterRotationRequested(address,address)`))              // VotingManagerUpgradeable
	SubmitterChosenTopic            = crypto.Keccak256Hash([]byte(`SubmitterChosen(address)`))                                 // VotingManagerUpgradeable
	TaskSubmittedTopic              = crypto.Keccak256Hash([]byte(`TaskSubmitted(uint256,bytes,address)`))                     // TaskManager
	TaskCompletedTopic              = crypto.Keccak256Hash([]byte(`TaskCompleted(uint256,address,uint256,bytes)`))             // TaskManager
	AddressRegisteredTopic          = crypto.Keccak256Hash([]byte(`AddressRegistered(address,uint256,uint8,uint256,address)`)) // AccountManagerContract
	ParticipantAddedTopic           = crypto.Keccak256Hash([]byte(`ParticipantAdded(address)`))                                // IParticipantManager
	ParticipantRemovedTopic         = crypto.Keccak256Hash([]byte(`ParticipantRemoved(address)`))                              // IParticipantManager
	DepositRecordedTopic            = crypto.Keccak256Hash([]byte(`DepositRecorded(address,uint256,uint256,bytes,bytes)`))     // DepositManagerContract
	WithdrawalRecordedTopic         = crypto.Keccak256Hash([]byte(`WithdrawalRecorded(address,uint256,uint256,bytes,bytes)`))  // DepositManagerContract

	WalletCreationRequestTopic = crypto.Keccak256Hash([]byte(`WalletCreationRequest(uint8,uint8,address,uint32,uint8,uint8)`))                                            // TopicPayloadContract
	DepositRequestTopic        = crypto.Keccak256Hash([]byte(`DepositRequest(uint8,uint8,string,uint64,uint8,uint32,uint64,string,string,string,uint8,uint8)`))           // TopicPayloadContract
	WithdrawalRequestTopic     = crypto.Keccak256Hash([]byte(`WithdrawalRequest(uint8,uint8,string,uint64,uint8,uint32,uint64,string,string,string,uint8,uint8,uint64)`)) // TopicPayloadContract
	WalletCreationResultTopic  = crypto.Keccak256Hash([]byte(`WalletCreationResult(uint8,bool,uint8,string,string)`))                                                     // TopicPayloadContract
	DepositResultTopic         = crypto.Keccak256Hash([]byte(`DepositResult(uint8,bool,uint8,string)`))                                                                   // TopicPayloadContract
	WithdrawalResultTopic      = crypto.Keccak256Hash([]byte(`WithdrawalResult(uint8,bool,uint8,string)`))                                                                // TopicPayloadContract
)

var topics = [][]common.Hash{
	{SubmitterRotationRequestedTopic},
	{SubmitterChosenTopic},
	{TaskSubmittedTopic},
	{TaskCompletedTopic},
	{AddressRegisteredTopic},
	{ParticipantRemovedTopic},
	{ParticipantAddedTopic},
	{DepositRecordedTopic},
	{WithdrawalRecordedTopic},
}

var (
	errNoEventSignature       = errors.New("no event signature")
	errEventSignatureMismatch = errors.New("event signature mismatch")
)

// UnpackLog unpacks a retrieved log into the provided output structure.
func UnpackLog(meta *bind.MetaData, out interface{}, event string, log types.Log) error {
	a, err := meta.GetAbi()
	if err != nil {
		return err
	}

	// Anonymous events are not supported.
	if len(log.Topics) == 0 {
		return errNoEventSignature
	}

	if log.Topics[0] != a.Events[event].ID {
		return errEventSignatureMismatch
	}

	if len(log.Data) > 0 {
		if err := a.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}

	var indexed abi.Arguments

	for _, arg := range a.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
