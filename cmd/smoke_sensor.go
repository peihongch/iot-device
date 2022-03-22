package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var smokeSensorCommand = &cobra.Command{
	Use:   "smoke",
	Short: "[mqtt] Start up smoke sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		smoke := sensor.NewSmokeSensor(data, remote, name)
		smoke.Start()
	},
}

func init() {
	smokeSensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")

	rootCmd.AddCommand(smokeSensorCommand)
}