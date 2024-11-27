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

	EncodeVerifyAndCall(operations []contracts.Operation, signature []byte) []byte
	GenerateVerifyTaskUnSignMsg(operations []contracts.Operation) (*big.Int, common.Hash, error)
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

func (l *Layer2Listener) GenerateVerifyTaskUnSignMsg(operations []contracts.Operation) (*big.Int, common.Hash, error) {
	nonce, err := l.contractVotingManager.TssNonce(nil)
	if err != nil {
		return nil, common.Hash{}, err
	}

	encodeData := contracts.EncodeOperation(nonce, operations)

	hash := crypto.Keccak256Hash(encodeData)
	hash = utils.PersonalMsgHash(hash)

	return nonce, hash, err
}

func (l *Layer2Listener) NextSubmitter() (common.Address, error) {
	return l.contractVotingManager.NextSubmitter(nil)
}

func (l *Layer2Listener) EncodeVerifyAndCall(operations []contracts.Operation, signature []byte) []byte {
	return contracts.EncodeFun(contracts.VotingManagerContractMetaData.ABI, "verifyAndCall", operations, signature)
}
