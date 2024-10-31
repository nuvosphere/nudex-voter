package abis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
)

var (
	VotingAddress      = common.HexToAddress(config.AppConfig.VotingContract)
	AccountAddress     = common.HexToAddress(config.AppConfig.AccountContract)
	OperationsAddress  = common.HexToAddress(config.AppConfig.OperationsContract)
	ParticipantAddress = common.HexToAddress(config.AppConfig.ParticipantContract)
)
