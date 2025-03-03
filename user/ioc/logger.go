package ioc

import (
	"tiktok_electric_business/pkg/logfile"

	"github.com/spf13/viper"
	"github.com/to404hanga/pkg404/logger"
	"go.uber.org/zap"
)

func InitLogger() logger.Logger {
	var cfg zap.Config
	mode := viper.GetString("log.mode")
	switch mode {
	case "dev":
		cfg = zap.NewDevelopmentConfig()
	case "prod":
		cfg = zap.NewProductionConfig()
	case "test":
		cfg = zap.NewDevelopmentConfig()
	default:
		mode = "dev"
		cfg = zap.NewDevelopmentConfig()
	}

	outputPaths := viper.GetStringSlice("log.outputPaths")
	errorOutputPaths := viper.GetStringSlice("log.errorOutputPaths")
	logfile.InitLogFilePath(outputPaths...)
	logfile.InitLogFilePath(errorOutputPaths...)
	cfg.OutputPaths = append(cfg.OutputPaths, outputPaths...)
	cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, errorOutputPaths...)

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
