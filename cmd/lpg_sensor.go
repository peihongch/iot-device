package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var lpgSensorCommand = &cobra.Command{
	Use:   "lpg",
	Short: "[mqtt] Start up lpg sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		lpg := sensor.NewLpgSensor(data, remote, name)
		lpg.Start()
	},
}

func init() {
	lpgSensorCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")
	rootCmd.AddCommand(lpgSensorCommand)
}
