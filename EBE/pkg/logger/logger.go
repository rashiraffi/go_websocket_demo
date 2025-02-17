package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(env string) (logger *zap.Logger, err error) {
	var zapConfig zap.Config
	if env == "PROD" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err = zapConfig.Build()
	if err != nil {
		log.Fatalln("Error building logger", err)
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}
