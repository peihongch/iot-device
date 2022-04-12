package air_alarm

import (
	"encoding/json"
	mqttclient "github.com/eclipse/paho.mqtt.golang"
	"github.com/peihongch/iot-device/pkg"
	"log"
	"time"
)

func NewAirAlarmClient(remote, name, token string) *AirAlarmClient {
	opts := mqttclient.NewClientOptions().AddBroker(remote).SetClientID(name)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(pkg.Handler)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetUsername(token)

	c := mqttclient.NewClient(opts)

	return &AirAlarmClient{
		topic:  name,
		remote: c,
	}
}

type AirAlarmClient struct {
	topic  string
	remote mqttclient.Client
}

func (ac AirAlarmClient) Start() {
	if token := ac.remote.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Println("air alarm client started")
	}
	defer ac.remote.Disconnect(250)

	if token := ac.remote.Subscribe(ac.topic, 0, func(client mqttclient.Client, message mqttclient.Message) {
		if err := ac.Execute(string(message.Payload())); err != nil {
			log.Fatalln("error execute air alarm command", err)
		}
	}); token.Error() != nil {
		log.Fatal("error subscribe edge platform", token.Error())
	} else {
		log.Printf("air alarm successfully subscribe %s", ac.topic)
	}

	for {
		time.Sleep(5 * time.Second)
	}
}

func (ac AirAlarmClient) Execute(cmd string) error {
	parsed := &AirAlarmCommand{}
	if err := json.Unmarshal([]byte(cmd), parsed); err != nil {
		return err
	} else {
		ExecCmd(parsed)
		return nil
	}
}
