package air_conditioner

import (
	"encoding/json"
	"fmt"
	mqttserver "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func NewAirConditionerBroker(name, topic, port string) *AirConditionerBroker {
	mqttserver.New()
	// Create the new MQTT Server.
	server := mqttserver.New()

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP(name, fmt.Sprintf("localhost:%s", port))

	// Add the listener to the server with default options (nil).
	err := server.AddListener(tcp, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &AirConditionerBroker{
		topic:  topic,
		port:   port,
		server: server,
	}
}

type AirConditionerBroker struct {
	topic  string
	port   string
	server *mqttserver.Server
}

func (ac AirConditionerBroker) Start() {
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

	log.Printf("air-conditioner mqttserver broker started: %v:%v\n", "0.0.0.0", ac.port)
	err := ac.server.Serve()
	if err != nil {
		log.Fatal(err)
	}

	<-done
	log.Println("interrupt signal caught")

	ac.server.Close()
	log.Println("air-conditioner mqttserver broker closed")
}

func (ac AirConditionerBroker) Execute(cmd string) error {
	parsed := &AirConditionerCommand{}
	if err := json.Unmarshal([]byte(cmd), parsed); err != nil {
		return err
	} else {
		ExecCmd(parsed)
		return nil
	}
}
