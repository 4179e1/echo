/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	pidFile string
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "The Echo Server",
	Long:  `The Echo Server for GRPC demo, with HTTP Gateway & Swagger UI support`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Global config
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/echo-server.yaml)")

	// server config
	rootCmd.PersistentFlags().String("server.pidfile", "/var/run/echo.pid", "the pid file")
	viper.BindPFlag("Server.PidFile", rootCmd.PersistentFlags().Lookup("server.pidfile"))

	// log config
	rootCmd.PersistentFlags().Bool("log.development", false, "is development config?")
	rootCmd.PersistentFlags().String("log.level", "debug", "log level")
	rootCmd.PersistentFlags().StringSlice("log.outputpaths", []string{"stdout"}, "output path")
	viper.BindPFlag("Log.Development", rootCmd.PersistentFlags().Lookup("log.development"))
	viper.BindPFlag("Log.Level", rootCmd.PersistentFlags().Lookup("log.level"))
	viper.BindPFlag("Log.OutputPaths", rootCmd.PersistentFlags().Lookup("log.outputpaths"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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
		viper.SetConfigName("echo-server")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("ECHO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	initLogger()

	sugar.Debug("Hello Echo Server...")
}

func initLogger() {
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

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	sugar = logger.Sugar()
}
