package htlc

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
)

type MsgCreateHTLC struct {
    Sender sdk.AccAddress
    Receiver string
    Amount sdk.Coin
    HashLock string
    TimeoutHeight clienttypes.Height
    TimeoutTimestamp uint64
    Memo string
}

type MsgClaimHTLC struct {
    Sender sdk.AccAddress
    SwapId string
    Secret string
}

type MsgRefundHTLC struct {
    Sender sdk.AccAddress
    SwapId string
}
