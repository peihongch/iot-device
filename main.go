package main

import "github.com/peihongch/iot-device/cmd"

func main() {
	cmd.Execute()
	//humidity := sensor.NewHumiditySensor("./data/devices/iot_telemetry_humidity.csv", "tcp://172.19.241.103:11883", "v1/devices/me/telemetry")
	//humidity.Start()
}
