package ioc

import (
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

	cfg.OutputPaths = append(cfg.OutputPaths, viper.GetStringSlice("log.outputPaths")...)
	cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, viper.GetStringSlice("log.errorOutputPaths")...)

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
