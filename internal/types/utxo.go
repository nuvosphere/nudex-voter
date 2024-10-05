package types

import "github.com/btcsuite/btcd/wire"

type BtcBlockExt struct {
	wire.MsgBlock

	BlockNumber uint64
}
