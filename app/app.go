package app

import (
    "github.com/cosmos/cosmos-sdk/baseapp"
    "github.com/cosmos/cosmos-sdk/codec"
    "cosmossdk.io/log"
    cosmosdb "github.com/cosmos/cosmos-db"
)

// App represents the MetalChained blockchain application
type App struct {
    BaseApp *baseapp.BaseApp
    cdc     *codec.LegacyAmino
}

// NewApp creates a new instance of the MetalChained app
func NewApp(logger log.Logger, db cosmosdb.DB) *App {
    cdc := codec.NewLegacyAmino()

    baseApp := baseapp.NewBaseApp("metalchained", logger, db, nil)

    return &App{
        BaseApp: baseApp,
        cdc:     cdc,
    }
}