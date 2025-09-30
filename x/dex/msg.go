package dex

import sdk "github.com/cosmos/cosmos-sdk/types"

type MsgCreatePool struct {
    Creator sdk.AccAddress
    TokenA string
    TokenB string
    AmountA sdk.Int
    AmountB sdk.Int
}
