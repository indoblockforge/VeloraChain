package htlc

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"

    sdk "github.com/cosmos/cosmos-sdk/types"
)

type HTLCStatus int

const (
    HTLCStatusPending HTLCStatus = iota
    HTLCStatusClaimed
    HTLCStatusRefunded
)

type HTLC struct {
    Id string
    Sender string
    Receiver string
    Amount sdk.Coin
    HashLock string
    Secret string
    Status HTLCStatus
}

type Keeper struct {
    store map[string]HTLC
    bankKeeper BankKeeper
}

type BankKeeper interface {
    SendCoinsFromAccountToModule(ctx interface{}, addr sdk.AccAddress, moduleName string, amt sdk.Coins) error
    SendCoinsFromModuleToAccount(ctx interface{}, moduleName string, addr sdk.AccAddress, amt sdk.Coins) error
}

func NewKeeper(bk BankKeeper) Keeper {
    return Keeper{
        store: make(map[string]HTLC),
        bankKeeper: bk,
    }
}

func sha256Hex(s string) string {
    h := sha256.Sum256([]byte(s))
    return hex.EncodeToString(h[:])
}

func (k *Keeper) CreateHTLC(sender sdk.AccAddress, receiver string, coin sdk.Coin, hashLock string) (string, error) {
    id := fmt.Sprintf("htlc-%d", len(k.store)+1)
    h := HTLC{
        Id: id,
        Sender: sender.String(),
        Receiver: receiver,
        Amount: coin,
        HashLock: hashLock,
        Status: HTLCStatusPending,
    }
    // move coins to module (omitted ctx)
    k.store[id] = h
    return id, nil
}

func (k *Keeper) ClaimHTLC(id string, secret string) error {
    h, ok := k.store[id]
    if !ok { return fmt.Errorf("not found") }
    if sha256Hex(secret) != h.HashLock { return fmt.Errorf("invalid secret") }
    h.Secret = secret
    h.Status = HTLCStatusClaimed
    k.store[id] = h
    return nil
}
