package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var gasSensorCommand = &cobra.Command{
	Use:   "gas",
	Short: "[mqtt] Start up gas (co, smoke and lpg) sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		co := sensor.NewGasSensor(data, remote, name, token)
		co.Start()
	},
}

func init() {
	gasSensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")
	gasSensorCommand.PersistentFlags().StringVar(&token, "token", "", "mqtt auth token, e.g. DHT22")

	rootCmd.AddCommand(gasSensorCommand)
}
