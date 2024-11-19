package util

import "go.uber.org/zap"

func LoggerInit() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()

	return logger, err
}
