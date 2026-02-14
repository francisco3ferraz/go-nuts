package logger

import (
	"go.uber.org/zap"
)

func New(environment string) *zap.Logger {
	if environment == "development" {
		cfg := zap.NewDevelopmentConfig()
		log, err := cfg.Build()
		if err != nil {
			panic(err)
		}

		return log
	}

	cfg := zap.NewProductionConfig()
	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return log
}
