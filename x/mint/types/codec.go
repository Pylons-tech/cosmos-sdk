package types

import (
	"github.com/Pylons-tech/cosmos-sdk/codec"
	cryptocodec "github.com/Pylons-tech/cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
