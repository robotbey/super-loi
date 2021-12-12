package main

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitLog() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	Log = logger.Sugar()
}
