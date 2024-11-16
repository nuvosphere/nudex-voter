package layer2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
)

type AccountManager interface {
	EncodeRegisterNewAddress(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address string) []byte
	GetAddressRecord(opts *bind.CallOpts, _user common.Address, _account *big.Int, _chain uint8, _index *big.Int) (string, error) // address
}

func (l *Layer2Listener) EncodeRegisterNewAddress(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address string) []byte {
	return contracts.EncodeFun(contracts.AccountManagerContractABI, "registerNewAddress", _user, _account, _chain, _index, _address)
}

func (l *Layer2Listener) GetAddressRecord(opts *bind.CallOpts, _user common.Address, _account *big.Int, _chain uint8, _index *big.Int) (string, error) {
	return l.accountManager.GetAddressRecord(nil, _user, _account, _chain, _index)
}
