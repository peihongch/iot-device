package executor

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

func NewAirConditioner(name string, topic string) *AirConditioner {
	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("emqx_test_client")

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	})
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)

	return &AirConditioner{
		topic:  topic,
		name:   name,
		remote: c,
	}
}

type AirConditioner struct {
	topic  string
	name   string
	remote mqtt.Client
}

func (ac AirConditioner) Start() {
	if token := ac.remote.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if err := ac.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func (ac AirConditioner) Execute() error {
	var err error
	ch := make(chan int, 1)

	// 订阅主题
	if token := ac.remote.Subscribe(ac.topic, 0, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("receive command: %s\n", string(message.Payload()))
	}); token.Wait() && token.Error() != nil {
		ch <- 1
		err = token.Error()
	}

	<-ch
	return err
}
