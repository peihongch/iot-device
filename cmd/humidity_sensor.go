package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var humiditySensorCommand = &cobra.Command{
	Use:   "humidity",
	Short: "[mqtt] Start up humidity sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		humidity := sensor.NewHumiditySensor(data, remote, name, token)
		humidity.Start()
	},
}

func init() {
	humiditySensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")
	humiditySensorCommand.PersistentFlags().StringVar(&token, "token", "", "mqtt auth token, e.g. DHT22")

	rootCmd.AddCommand(humiditySensorCommand)
}
