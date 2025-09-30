package gov

import (
    "encoding/json"

    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    abci "github.com/tendermint/tendermint/abci/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    gov "github.com/cosmos/cosmos-sdk/x/gov"
)

var (
    _ module.AppModule      = AppModule{}
    _ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
    gov.AppModuleBasic
}

type AppModule struct {
    gov.AppModule
}

func NewAppModule(cdc codec.Codec, keeper gov.Keeper) AppModule {
    return AppModule{gov.NewAppModule(cdc, keeper)}
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
    am.AppModule.BeginBlock(ctx, req)
}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
    return am.AppModule.EndBlock(ctx, req)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    return gov.AppModuleBasic{}.DefaultGenesis(cdc)
}
