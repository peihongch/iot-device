package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var tempSensorCommand = &cobra.Command{
	Use:   "temp",
	Short: "[mqtt] Start up temperature sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		temp := sensor.NewTempSensor(data, remote, name)
		temp.Start()
	},
}

func init() {
	tempSensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")

	rootCmd.AddCommand(tempSensorCommand)
}
