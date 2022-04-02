package cmd

import (
	"github.com/peihongch/iot-device/pkg/sensor"
	"github.com/spf13/cobra"
)

var (
	host = ""
	path = ""
)

var motionSensorCommand = &cobra.Command{
	Use:   "motion",
	Short: "[coap] Start up motion sensor device",
	Run: func(cmd *cobra.Command, args []string) {
		motion := sensor.NewMotionSensor(data, &sensor.CoapOpts{
			Host: host,
			Port: port,
			Path: path,
		}, name)
		motion.Start()
	},
}

func init() {
	motionSensorCommand.PersistentFlags().StringVar(&host, "host", "", "coap server host")
	motionSensorCommand.PersistentFlags().StringVar(&port, "port", "", "coap server port")
	motionSensorCommand.PersistentFlags().StringVar(&path, "path", "", "coap server gateway path")

	rootCmd.AddCommand(motionSensorCommand)
}
