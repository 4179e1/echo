package common

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BindRootFlag(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().Bool("log.development", false, "is development config?")
	rootCmd.PersistentFlags().String("log.level", "info", "log level")
	rootCmd.PersistentFlags().StringSlice("log.outputpaths", []string{"stdout"}, "output path")
	viper.BindPFlag("Log.Development", rootCmd.PersistentFlags().Lookup("log.development"))
	viper.BindPFlag("Log.Level", rootCmd.PersistentFlags().Lookup("log.level"))
	viper.BindPFlag("Log.OutputPaths", rootCmd.PersistentFlags().Lookup("log.outputpaths"))
}
