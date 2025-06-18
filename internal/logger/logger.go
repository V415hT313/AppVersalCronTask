package logger

import (
	"go.uber.org/zap"
)

func Init() {

	logger, _ := zap.NewDevelopment()

	zap.ReplaceGlobals(logger)
	zap.L().Info("Zap logger initialized")
}

func CloseLogger() {
    zap.L().Sync()
}