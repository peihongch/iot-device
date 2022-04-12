package cmd

import (
	air_conditioner "github.com/peihongch/iot-device/pkg/executor/air-conditioner"
	"github.com/spf13/cobra"
	"log"
)

var airConditionerCommand = &cobra.Command{
	Use:   "air-conditioner",
	Short: "[mqtt] Start up air conditioner device",
	Run: func(cmd *cobra.Command, args []string) {
		switch mode {
		case "client":
			airConditioner := air_conditioner.NewAirConditionerClient(remote, name, token)
			airConditioner.Start()
		case "broker":
			airConditioner := air_conditioner.NewAirConditionerBroker(name, topic, port)
			airConditioner.Start()
		default:
			log.Fatalf("bad mqtt mode: %v, only support either `client` or `broker`", mode)
		}
	},
}

func init() {
	airConditionerCommand.PersistentFlags().StringVar(&mode, "mode", "client", "mqtt mode to start, client or broker")

	airConditionerCommand.PersistentFlags().StringVar(&topic, "topic", "", "mqtt topic to subscribe (for client mode) or publish (for server mode)")
	airConditionerCommand.PersistentFlags().StringVar(&remote, "remote", "", "(for client mode only) mqtt server url, e.g. tcp://broker.emqx.io:1883")
	airConditionerCommand.PersistentFlags().StringVar(&token, "token", "", "(for client mode only) mqtt auth token, e.g. DHT22")
	airConditionerCommand.PersistentFlags().StringVar(&port, "port", "1883", "(for broker mode only) mqtt broker port url, default 1883")

	rootCmd.AddCommand(airConditionerCommand)
}
