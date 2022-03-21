package pkg

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Device 传感器设备接口
type Device interface {
	// Start 启动设备守护进程
	Start()
	// Collect 传感器设备采集环境数据
	Collect() error
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
