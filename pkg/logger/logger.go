package logger

import (
	"context"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(lc fx.Lifecycle) (*zap.Logger, error) {

	// Rotating file writer
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/app.log", // path to log file
		MaxSize:    100,              // MB per file
		MaxBackups: 10,               // number of old files to keep
		MaxAge:     30,               // days
		Compress:   true,             // gzip old logs
	})

	consoleWriter := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, fileWriter, zap.InfoLevel),
		zapcore.NewCore(jsonEncoder, consoleWriter, zap.InfoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	// Proper shutdown
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger, nil

	// cfg := zap.NewProductionConfig()
	// cfg.Encoding = "json"
	// return cfg.Build()
}
