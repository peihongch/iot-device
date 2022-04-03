package cmd

import (
	"github.com/peihongch/iot-device/pkg/executor"
	"github.com/spf13/cobra"
)

var airAlarmCommand = &cobra.Command{
	Use:   "air-alarm",
	Short: "[mqtt] Start up air alarm device",
	Run: func(cmd *cobra.Command, args []string) {
		airAlarm := executor.NewAirAlarm(name, topic, port)
		airAlarm.Start()
	},
}

func init() {
	airAlarmCommand.PersistentFlags().StringVar(&port, "port", "1884", "mqtt broker port url, default 1884")
	airAlarmCommand.PersistentFlags().StringVar(&topic, "topic", "", "mqtt topic to subscribe")

	rootCmd.AddCommand(airAlarmCommand)
}
