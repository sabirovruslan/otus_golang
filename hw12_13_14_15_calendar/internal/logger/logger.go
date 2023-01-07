package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.String() + "]")
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05"))
}

type Logger struct {
	logger *zap.Logger
}

func New(level string) *Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.LevelKey = level
	cfg.EncoderConfig.EncodeLevel = customLevelEncoder
	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	logger, _ := cfg.Build()
	return &Logger{
		logger: logger,
	}
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

// TODO
