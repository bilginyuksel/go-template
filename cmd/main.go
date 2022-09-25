package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	conf := readConfig()
	logger.Info("application started..", zap.Any("conf", conf))
}
