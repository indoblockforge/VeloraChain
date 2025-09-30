package dex

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
    StoreKey = "x_dex"
)

type Pool struct {
    PoolId     uint64
    Creator    string
    TokenA     string
    TokenB     string
    ReserveA   sdk.Int
    ReserveB   sdk.Int
    TotalShares sdk.Int
}
