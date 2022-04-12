package air_alarm

import (
	"encoding/json"
	mqttclient "github.com/eclipse/paho.mqtt.golang"
	"github.com/peihongch/iot-device/pkg"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	if token := ac.remote.Subscribe(ac.topic, 0, func(client mqttclient.Client, message mqttclient.Message) {

	}); token.Error() != nil {
		log.Fatal("error subscribe edge platform", token.Error())
	}

	<-done
	log.Println("interrupt signal caught")

	ac.remote.Unsubscribe(ac.topic)
	ac.remote.Disconnect(0)
	log.Println("air-alarm mqtt broker closed")
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
