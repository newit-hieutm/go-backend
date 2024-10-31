package loggers

import (
	"fmt"
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

	var core zapcore.Core
	levelLog := zap.InfoLevel
	var productionCfg zapcore.EncoderConfig = zap.NewProductionEncoderConfig()

	env := os.Getenv("GO_ENV")
	isLocal := env == "development" || env == ""

	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fmt.Println("ENVIRONMENT IS", env)

	if isLocal {
		// Local environment: log to both file and stdout
		levelLog = zap.DebugLevel
		productionCfg = zap.NewDevelopmentEncoderConfig()

		file := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    10, // megabytes
			MaxBackups: 3,
			MaxAge:     7, // days
		})

		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
		fileEncoder := zapcore.NewJSONEncoder(productionCfg)

		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, file, zap.NewAtomicLevelAt(levelLog)),
			zapcore.NewCore(consoleEncoder, stdout, zap.NewAtomicLevelAt(levelLog)),
		)
	} else {
		// Production: log only to stdout
		encoder := zapcore.NewJSONEncoder(productionCfg)
		core = zapcore.NewCore(encoder, stdout, zap.NewAtomicLevelAt(levelLog))
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return logger
}
