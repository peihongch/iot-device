package sensor

import (
	"encoding/csv"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/peihongch/iot-device/pkg"
	"io"
	"log"
	"os"
	"time"
)

// NewLightSensor 实例化亮度传感器
//  source 数据源
//  remote 数据发送的远端目的平台
func NewLightSensor(source string, remote string, name string, token string) *LightSensor {
	opts := mqtt.NewClientOptions().AddBroker(remote).SetClientID(name)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(pkg.Handler)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetUsername(token)

	c := mqtt.NewClient(opts)

	fs, err := os.Open(source)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	r := csv.NewReader(fs)
	// 丢弃首行
	_, err = r.Read()
	if err != nil {
		log.Fatalf("error discard csv header, err is %+v", err)
		return nil
	}

	return &LightSensor{
		topic:  name,
		remote: c,
		source: r,
	}
}

// LightSensor 亮度传感器
type LightSensor struct {
	topic  string
	remote mqtt.Client
	source *csv.Reader
}

type LightData struct {
	Timestamp string `json:"ts"`
	Device    string `json:"device"`
	Light     string `json:"light"`
}

func (t LightSensor) Collect() error {
	row, err := t.source.Read()
	if err != nil {
		return err
	}

	if data, err := json.Marshal(&LightData{
		Timestamp: row[0],
		Device:    row[1],
		Light:     row[2],
	}); err != nil {
		log.Println("error when marshaling", err)
		return err
	} else {
		t.remote.Publish(t.topic, 0, false, data)
		log.Println(string(data))
		return nil
	}
}

func (t LightSensor) Start() {
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
