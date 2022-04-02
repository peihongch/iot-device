package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var coSensorCommand = &cobra.Command{
	Use:   "co",
	Short: "[mqtt] Start up co sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		co := sensor.NewCoSensor(data, remote, name, token)
		co.Start()
	},
}

func init() {
	coSensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")
	coSensorCommand.PersistentFlags().StringVar(&token, "token", "", "mqtt auth token, e.g. DHT22")

	rootCmd.AddCommand(coSensorCommand)
}
