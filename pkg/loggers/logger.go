package loggers

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var instanceLogger *zap.Logger

func InitLogger() *zap.Logger {
	if instanceLogger == nil {
		instanceLogger = CreateLogger()
	}

	return instanceLogger
}

func CreateLogger() *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	levelLog := zap.InfoLevel

	var productionCfg zapcore.EncoderConfig = zap.NewProductionEncoderConfig()

	if os.Getenv("ENV") == "local" || os.Getenv("ENV") == "" {
		levelLog = zap.DebugLevel
		productionCfg = zap.NewDevelopmentEncoderConfig()
	}

	level := zap.NewAtomicLevelAt(levelLog)

	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)

	// Add WithCaller() to include caller information in logs
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return logger
}
