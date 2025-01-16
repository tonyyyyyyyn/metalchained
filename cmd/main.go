package main

import (
    "cosmossdk.io/log"
    cosmosdb "github.com/cosmos/cosmos-db"
    "github.com/tonynovatsky/metalchained/app"
)

func main() {
    logger := log.NewNopLogger()
    database := cosmosdb.NewMemDB()

    metalChainedApp := app.NewApp(logger, database)

    logger.Info("MetalChained App initialized!", "app", metalChainedApp)
}