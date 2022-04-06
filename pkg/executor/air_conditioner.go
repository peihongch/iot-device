package executor

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/peihongch/iot-device/pkg"
	"log"
)

func NewAirConditioner(name, topic, port string) *AirConditioner {
	mqtt.New()
	// Create the new MQTT Server.
	server := mqtt.New()

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP(name, fmt.Sprintf("localhost:%s", port))

	// Add the listener to the server with default options (nil).
	err := server.AddListener(tcp, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &AirConditioner{
		topic:  topic,
		port:   port,
		server: server,
	}
}

type AirConditioner struct {
	topic  string
	port   string
	server *mqtt.Server
}

type AirConditionerCommand struct {
	Mode   string `json:"mode"`
	Target int    `json:"target"`
}

func (ac AirConditioner) Start() {
	// Start the broker. Serve() is blocking - see examples folder
	// for usage ideas.
	ac.server.Events.OnMessage = func(client events.Client, packet events.Packet) (pk events.Packet, err error) {
		err = ac.Execute(string(packet.Payload))
		return
	}

	log.Printf("mqtt broker started: %v:%v\n", "0.0.0.0", ac.port)
	err := ac.server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (ac AirConditioner) Execute(cmd string) error {
	parsed := &AirConditionerCommand{}
	err := json.Unmarshal([]byte(cmd), parsed)
	if err != nil {
		return err
	}

	switch parsed.Mode {
	case "hot":
		fmt.Printf("空调【模式：%s】【温度：%v℃】\n", pkg.SprintRed("制热"), pkg.SprintCyan(fmt.Sprintf("%v", parsed.Target)))
	case "cold":
		fmt.Printf("空调【模式：%s】【温度：%v℃】\n", pkg.SprintBlue("制冷"), pkg.SprintCyan(fmt.Sprintf("%v", parsed.Target)))
	case "dry":
		fmt.Printf("空调【模式：%s】\n", pkg.SprintGreen("除湿"))
	}

	return nil
}
