package main

import (
    "log"

    "github.com/velorachain/core/app"
)

func main() {
    // Minimal entry — real server uses Ethermint server wrapper.
    log.Println("VeloraChain node (scaffold) — replace with full server start")
    _ = app.NewVeloraAppBuilder(nil)
}
