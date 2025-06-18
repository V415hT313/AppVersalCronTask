package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/V415hT313/AppVersalCronTask/internal/cron"
	"github.com/V415hT313/AppVersalCronTask/internal/logger"
	"github.com/V415hT313/AppVersalCronTask/internal/server"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.CloseLogger() 

	zap.L().Info("Starting the application...")

	go func() {
		zap.L().Info("Starting gRPC server...")
		server.StartGRPCServer()
	}()

	go func() {
		zap.L().Info("Starting cron job...")
		cron.StartCronJob()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Application is running. Press CTRL+C to exit.")
	
}
