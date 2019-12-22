package common

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(logger **zap.Logger, sugar **zap.SugaredLogger) {
	var level zapcore.Level
	// TODO: valid Log.Level input
	level.Set(viper.GetString("Log.Level"))
	atom := zap.NewAtomicLevelAt(level)

	isDevelopment := viper.GetBool("Log.Development")
	var cfg zap.Config
	if isDevelopment {
		cfg = zap.NewDevelopmentConfig()

	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.Level = atom
	cfg.OutputPaths = viper.GetStringSlice("Log.OutputPaths")

	var err error
	*logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	*sugar = (*logger).Sugar()

	fmt.Println("=============================")
}
