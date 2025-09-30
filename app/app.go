package app

import (
    "io"

    "github.com/tendermint/tendermint/libs/log"
)

// NewVeloraAppBuilder returns a builder function (scaffold)
func NewVeloraAppBuilder(cfg interface{}) func(logger log.Logger, db interface{}, traceStore io.Writer) interface{} {
    return func(logger log.Logger, db interface{}, traceStore io.Writer) interface{} {
        logger.Info("building VeloraApp (scaffold)")
        return nil
    }
}
