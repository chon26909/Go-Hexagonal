package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, field ...zap.Field) {
	log.Info(message, field...)
}

func Debug(message string, field ...zap.Field) {
	log.Debug(message, field...)
}

func Warn(message string, field ...zap.Field) {
	log.Warn(message, field...)
}

func Error(message string, field ...zap.Field) {
	log.Error(message, field...)
}
