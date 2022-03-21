package pkg

import (
	"encoding/csv"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io"
	"log"
	"os"
	"time"
)

// NewCoSensor 实例化CO传感器
//  source 数据源
//  remote 数据发送的远端目的平台
func NewCoSensor(source string, remote string, name string) *CoSensor {
	opts := mqtt.NewClientOptions().AddBroker(remote).SetClientID(name)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)

	fs, err := os.Open(source)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	r := csv.NewReader(fs)

	return &CoSensor{
		topic:  name,
		remote: c,
		source: r,
	}
}

// CoSensor CO传感器
type CoSensor struct {
	topic  string
	remote mqtt.Client
	source *csv.Reader
}

type CoData struct {
	Timestamp string `json:"ts"`
	Device    string `json:"device"`
	CO        string `json:"co"`
}

func (t CoSensor) Collect() error {
	row, err := t.source.Read()
	if err != nil {
		return err
	}

	if data, err := json.Marshal(&CoData{
		Timestamp: row[0],
		Device:    row[1],
		CO:        row[2],
	}); err != nil {
		log.Println("error when marshaling", err)
		return err
	} else {
		t.remote.Publish(t.topic, 0, false, data)
		return nil
	}
}

func (t CoSensor) Start() {
	if token := t.remote.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer t.remote.Disconnect(250)

	for range time.Tick(5 * time.Second) {
		err := t.Collect()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			log.Println("no data anymore")
		}
	}
}
