package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(cfgFile string, defaultCfgFile string, logger **zap.Logger, sugar **zap.SugaredLogger) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "echo-server" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(defaultCfgFile)
	}
	fmt.Println(defaultCfgFile)

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("ECHO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatalf(fmt.Sprintf("Error loading configuration file(%s): %s", viper.ConfigFileUsed(), err.Error()))
	}

	InitLogger(logger, sugar)

}
