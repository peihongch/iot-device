package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var thSensorCommand = &cobra.Command{
	Use:   "th",
	Short: "[mqtt] Start up th (temperature and humidity) sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		humidity := sensor.NewTempHumiditySensor(data, remote, name, token)
		humidity.Start()
	},
}

func init() {
	thSensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")
	thSensorCommand.PersistentFlags().StringVar(&token, "token", "", "mqtt auth token, e.g. DHT22")

	rootCmd.AddCommand(thSensorCommand)
}
