package main

import (
    "github.com/V415hT313/AppVersalCronTask/internal/cron"
    "github.com/V415hT313/AppVersalCronTask/internal/logger"
    "github.com/V415hT313/AppVersalCronTask/internal/server"
)

func main() {
    logger.Init()
    go server.StartGRPCServer()
    go cron.StartCronJob()

    select {}
}
