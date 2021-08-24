package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Pylons-tech/cosmos-sdk/codec"
	codectypes "github.com/Pylons-tech/cosmos-sdk/codec/types"
	"github.com/Pylons-tech/cosmos-sdk/std"
	"github.com/Pylons-tech/cosmos-sdk/testutil/testdata"
	sdk "github.com/Pylons-tech/cosmos-sdk/types"
	"github.com/Pylons-tech/cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
