package cmd

import (
	"time"

	"github.com/lazyshot/vatsim-data/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var pollCmd = &cobra.Command{
	Use: "poll",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New()

		for range time.NewTicker(viper.GetDuration("interval")).C {
			data, err := c.PullData()
			if err != nil {
				zap.S().Errorw("failed to fetch vatsim data",
					"err", err,
				)
				continue
			}

			zap.S().Debugw("data returned",
				"data", data,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(pollCmd)

	pollCmd.Flags().Duration("interval", 15*time.Second, "Time between pulls from the VATSIM data server")
	viper.BindPFlags(pollCmd.Flags())
}
