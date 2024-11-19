package util

import (
	"go.uber.org/zap"
)

func LoggerInit(config Configuration) (*zap.Logger, error) {

	var logger *zap.Logger
	var err error
	if config.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewDevelopment()

	}

	return logger, err
}
