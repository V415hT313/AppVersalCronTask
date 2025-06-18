package main

import (
    "github.com/V415hT313/AppVersalCronTask/internal/logger"
    "github.com/V415hT313/AppVersalCronTask/internal/server"
    "github.com/V415hT313/AppVersalCronTask/internal/cron"
)

func main() {
    logger.Init()
    go server.StartGRPCServer()
    go cron.StartCronJob()

    // Keep the main function running
    select {}
}
