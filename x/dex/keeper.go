package dex

import (
    "fmt"

    sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
    bankKeeper BankKeeper
    store map[string]Pool
    nextPoolId uint64
}

type BankKeeper interface {
    MintCoins(ctx interface{}, moduleName string, amt sdk.Coins) error
    SendCoinsFromModuleToAccount(ctx interface{}, moduleName string, addr sdk.AccAddress, amt sdk.Coins) error
    SendCoinsFromAccountToModule(ctx interface{}, addr sdk.AccAddress, moduleName string, amt sdk.Coins) error
}

func NewKeeper(bk BankKeeper) Keeper {
    return Keeper{
        bankKeeper: bk,
        store: make(map[string]Pool),
        nextPoolId: 1,
    }
}

func (k *Keeper) CreatePool(ctx interface{}, creator sdk.AccAddress, tokenA, tokenB string, amtA, amtB sdk.Int) (uint64, error) {
    id := k.nextPoolId
    k.nextPoolId++
    p := Pool{
        PoolId: id,
        Creator: creator.String(),
        TokenA: tokenA,
        TokenB: tokenB,
        ReserveA: amtA,
        ReserveB: amtB,
        TotalShares: sdk.NewInt(0),
    }
    k.store[fmt.Sprintf("%d", id)] = p
    return id, nil
}
