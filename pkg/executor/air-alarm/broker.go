package air_alarm

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func NewAirAlarmBroker(name, topic, port string) *AirAlarmBroker {
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

	return &AirAlarmBroker{
		topic:  topic,
		port:   port,
		server: server,
	}
}

type AirAlarmBroker struct {
	topic  string
	port   string
	server *mqtt.Server
}

func (ac AirAlarmBroker) Start() {
	// Start the broker. Serve() is blocking - see examples folder
	// for usage ideas.
	ac.server.Events.OnMessage = func(client events.Client, packet events.Packet) (pk events.Packet, err error) {
		err = ac.Execute(string(packet.Payload))
		return
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	log.Printf("air-alarm mqtt broker started: %v:%v\n", "0.0.0.0", ac.port)
	err := ac.server.Serve()
	if err != nil {
		log.Fatal(err)
	}

	<-done
	log.Println("interrupt signal caught")

	ac.server.Close()
	log.Println("air-alarm mqtt broker closed")
}

func (ac AirAlarmBroker) Execute(cmd string) error {
	parsed := &AirAlarmCommand{}
	if err := json.Unmarshal([]byte(cmd), parsed); err != nil {
		return err
	} else {
		ExecCmd(parsed)
		return nil
	}
}
