package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/nats-io/nats.go"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var log *zap.Logger

var rootCmd = &cobra.Command{
	Use: "vatsim-data",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		log, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

		zap.ReplaceGlobals(log)
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		log.Sync()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("nats-url", nats.DefaultURL, "url to nats server")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}
