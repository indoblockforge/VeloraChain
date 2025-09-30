package gov

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type Keeper struct {
    govkeeper.Keeper
}

func (k Keeper) SubmitTextProposal(ctx sdk.Context, title, description string, proposer sdk.AccAddress) (govtypes.Proposal, error) {
    content := govtypes.NewTextProposal(title, description)
    return k.Keeper.SubmitProposal(ctx, content)
}
