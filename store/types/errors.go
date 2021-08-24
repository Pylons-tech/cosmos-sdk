package types

import (
	sdkerrors "github.com/Pylons-tech/cosmos-sdk/types/errors"
)

const StoreCodespace = "store"

var (
	ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
)
