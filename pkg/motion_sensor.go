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

// NewMotionSensor 实例化动作传感器
//  source 数据源
//  remote 数据发送的远端目的平台
func NewMotionSensor(source string, remote string, name string) *MotionSensor {
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

	return &MotionSensor{
		topic:  name,
		remote: c,
		source: r,
	}
}

// MotionSensor 动作传感器
type MotionSensor struct {
	topic  string
	remote mqtt.Client
	source *csv.Reader
}

type MotionData struct {
	Timestamp string `json:"ts"`
	Device    string `json:"device"`
	Motion    string `json:"motion"`
}

func (t MotionSensor) Collect() error {
	row, err := t.source.Read()
	if err != nil {
		return err
	}

	if data, err := json.Marshal(&MotionData{
		Timestamp: row[0],
		Device:    row[1],
		Motion:    row[2],
	}); err != nil {
		log.Println("error when marshaling", err)
		return err
	} else {
		t.remote.Publish(t.topic, 0, false, data)
		return nil
	}
}

func (t MotionSensor) Start() {
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
