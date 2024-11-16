package layer2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

type ContractVotingManager interface {
	TssNonce() (*big.Int, error)
	Proposer() (common.Address, error)
	NextSubmitter() (common.Address, error)
	TssSigner() (common.Address, error)
	LastSubmissionTime() (*big.Int, error)
	ForcedRotationWindow() (*big.Int, error)
	TaskCompletionThreshold() (*big.Int, error)

	EncodeVerifyAndCall(_target common.Address, _data []byte, _signature []byte) []byte
	GenerateVerifyTaskUnSignMsg(contractAddress common.Address, calldata []byte, taskID *big.Int) (common.Hash, error)
}

func (l *Layer2Listener) TssSigner() (common.Address, error) {
	return l.contractVotingManager.TssSigner(nil)
}

func (l *Layer2Listener) LastSubmissionTime() (*big.Int, error) {
	return l.contractVotingManager.LastSubmissionTime(nil)
}

func (l *Layer2Listener) ForcedRotationWindow() (*big.Int, error) {
	return l.contractVotingManager.ForcedRotationWindow(nil)
}

func (l *Layer2Listener) TaskCompletionThreshold() (*big.Int, error) {
	return l.contractVotingManager.TaskCompletionThreshold(nil)
}

func (l *Layer2Listener) TssNonce() (*big.Int, error) {
	return l.contractVotingManager.TssNonce(nil)
}

func (l *Layer2Listener) ContractVotingManager() *contracts.VotingManagerContract {
	return l.contractVotingManager
}

func (l *Layer2Listener) Proposer() (common.Address, error) {
	return l.contractVotingManager.NextSubmitter(nil)
}

func (l *Layer2Listener) GenerateVerifyTaskUnSignMsg(contractAddress common.Address, calldata []byte, taskID *big.Int) (common.Hash, error) {
	nonce, err := l.contractVotingManager.TssNonce(nil)
	if err != nil {
		return common.Hash{}, nil
	}

	nonce.Add(nonce, big.NewInt(1))

	encodeData, err := utils.AbiEncodePacked(nonce, contractAddress, calldata, taskID)
	if err != nil {
		return common.Hash{}, nil
	}

	return crypto.Keccak256Hash(encodeData), nil
}

func (l *Layer2Listener) NextSubmitter() (common.Address, error) {
	return l.contractVotingManager.NextSubmitter(nil)
}

func (l *Layer2Listener) EncodeVerifyAndCall(_target common.Address, _data []byte, _signature []byte) []byte {
	return contracts.EncodeFun(contracts.VotingManagerContractABI, "verifyAndCall", _target, _data, _signature)
}
