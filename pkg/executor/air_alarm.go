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

func NewAirAlarm(name, topic, port string) *AirAlarm {
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

	return &AirAlarm{
		topic:  topic,
		port:   port,
		server: server,
	}
}

type AirAlarm struct {
	topic  string
	port   string
	server *mqtt.Server
}

type AirAlarmCommand struct {
	Timestamp int     `json:"ts"`
	Device    string  `json:"device"`
	Co        float64 `json:"co"`
	Threshold float64 `json:"threshold"`
}

func (ac AirAlarm) Start() {
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

func (ac AirAlarm) Execute(cmd string) error {
	parsed := &AirAlarmCommand{}
	err := json.Unmarshal([]byte(cmd), parsed)
	if err != nil {
		return err
	}

	if parsed.Co >= parsed.Threshold {
		fmt.Println(pkg.SprintRed(fmt.Sprintf("!!!WARNING!!! 一氧化碳浓度到达 %v，危险阈值为 %v!", parsed.Co, parsed.Threshold)))
	}

	return nil
}
