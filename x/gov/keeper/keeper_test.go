package keeper

import (
    "testing"

    sdk "github.com/cosmos/cosmos-sdk/types"
    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
    "github.com/stretchr/testify/require"
)

func TestSubmitTextProposal(t *testing.T) {
    ctx, k := setupKeeper(t)
    addr := sdk.AccAddress([]byte("addr1---------------"))

    proposal, err := k.SubmitTextProposal(ctx, "Test Title", "Test Description", addr)
    require.NoError(t, err)
    require.Equal(t, govtypes.StatusDepositPeriod, proposal.Status)
}
