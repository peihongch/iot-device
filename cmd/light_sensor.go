package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var lightSensorCommand = &cobra.Command{
	Use:   "light",
	Short: "[mqtt] Start up light sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		light := sensor.NewLightSensor(data, remote, name)
		light.Start()
	},
}

func init() {
	lightSensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")

	rootCmd.AddCommand(lightSensorCommand)
}
