package cmd

import (
	"github.com/peihongch/iot-device/pkg/executor"
	"github.com/spf13/cobra"
)

var airConditionerCommand = &cobra.Command{
	Use:   "air-conditioner",
	Short: "[mqtt] Start up air conditioner device",
	Run: func(cmd *cobra.Command, args []string) {
		airConditioner := executor.NewAirConditioner(name, topic, remote)
		airConditioner.Start()
	},
}

func init() {
	airConditionerCommand.PersistentFlags().StringVar(&remote, "remote", "", "mqtt server url, e.g. tcp://broker.emqx.io:1883")
	airConditionerCommand.PersistentFlags().StringVar(&topic, "topic", "", "mqtt topic to subscribe")

	rootCmd.AddCommand(airConditionerCommand)
}
