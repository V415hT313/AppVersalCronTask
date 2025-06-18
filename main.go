package main

import (
	"github.com/V415hT313/AppVersalCronTask/internal/cron"
	"github.com/V415hT313/AppVersalCronTask/internal/logger"
	"github.com/V415hT313/AppVersalCronTask/internal/server"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	zap.L().Info("Starting the application...")

	go func() {
		zap.L().Info("Starting gRPC server...")
		server.StartGRPCServer()
	}()

	go func() {
		zap.L().Info("Starting cron job...")
		cron.StartCronJob()
	}()

	zap.L().Info("Application is running. Press CTRL+C to exit.")
	select {}
}
