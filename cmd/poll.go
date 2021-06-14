package cmd

import (
	"encoding/json"
	"time"

	"github.com/lazyshot/vatsim-data/pkg/client"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var pollCmd = &cobra.Command{
	Use: "poll",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New()
		nc, err := nats.Connect(viper.GetString("nats-url"))
		if err != nil {
			zap.S().Fatalw("fail to connect to nats", "err", err)
		}

		js, err := nc.JetStream()
		if err != nil {
			zap.S().Fatalw("fail to connect to jetstream", "err", err)
		}

		for range time.NewTicker(viper.GetDuration("interval")).C {
			data, err := c.PullData()
			if err != nil {
				zap.S().Errorw("failed to fetch vatsim data",
					"err", err,
				)
				continue
			}

			for _, p := range data.Pilots {
				b, err := json.Marshal(p)
				if err != nil {
					zap.S().Errorw("failed to marshal pilot data",
						"err", err,
					)
					continue
				}

				js.Publish("pilots", b)
			}

			for _, c := range data.Controllers {
				b, err := json.Marshal(c)
				if err != nil {
					zap.S().Errorw("failed to marshal pilot data",
						"err", err,
					)
					continue
				}

				js.Publish("controllers", b)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pollCmd)

	pollCmd.Flags().Duration("interval", 15*time.Second, "Time between pulls from the VATSIM data server")
	viper.BindPFlags(pollCmd.Flags())
}
