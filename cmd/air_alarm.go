package cmd

import (
	air_alarm "github.com/peihongch/iot-device/pkg/executor/air-alarm"
	"github.com/spf13/cobra"
	"log"
)

var airAlarmCommand = &cobra.Command{
	Use:   "air-alarm",
	Short: "[mqtt] Start up air alarm device",
	Run: func(cmd *cobra.Command, args []string) {
		switch mode {
		case "client":
			airAlarm := air_alarm.NewAirAlarmClient(remote, name, token)
			airAlarm.Start()
		case "broker":
			airAlarm := air_alarm.NewAirAlarmBroker(name, topic, port)
			airAlarm.Start()
		default:
			log.Fatalf("bad mqtt mode: %v, only support either `client` or `broker`", mode)
		}
	},
}

func init() {
	airAlarmCommand.PersistentFlags().StringVar(&mode, "mode", "client", "mqtt mode to start, client or broker")

	airAlarmCommand.PersistentFlags().StringVar(&topic, "topic", "", "mqtt topic to subscribe (for client mode) or publish (for server mode)")
	airAlarmCommand.PersistentFlags().StringVar(&remote, "remote", "", "(for client mode only) mqtt server url, e.g. tcp://broker.emqx.io:1883")
	airAlarmCommand.PersistentFlags().StringVar(&token, "token", "", "(for client mode only) mqtt auth token, e.g. DHT22")
	airAlarmCommand.PersistentFlags().StringVar(&port, "port", "1883", "(for broker mode only) mqtt broker port url, default 1883")

	rootCmd.AddCommand(airAlarmCommand)
}
