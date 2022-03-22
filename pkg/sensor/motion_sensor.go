package sensor

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/udp"
	"github.com/plgd-dev/go-coap/v2/udp/client"
	"io"
	"log"
	"os"
	"time"
)

// NewMotionSensor 实例化动作传感器
//  source 数据源
//  remote 数据发送的远端目的平台
func NewMotionSensor(source string, opts *CoapOpts, name string) *MotionSensor {
	fs, err := os.Open(source)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	r := csv.NewReader(fs)

	return &MotionSensor{
		topic:  name,
		remote: fmt.Sprintf("%s:%s", opts.Host, opts.Port),
		path:   opts.Path,
		source: r,
	}
}

// MotionSensor 动作传感器
type MotionSensor struct {
	ctx    context.Context
	topic  string
	path   string
	remote string
	co     *client.ClientConn
	source *csv.Reader
}

type CoapOpts struct {
	Host string
	Port string
	Path string
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
		if msg, err := t.co.Post(t.ctx, t.path, message.TextPlain, bytes.NewReader(data)); err != nil {
			log.Println("error when sending coap post message", err)
		} else {
			log.Println("coap resp", msg)
		}
		return nil
	}
}

func (t *MotionSensor) Start() {
	co, err := udp.Dial(t.remote)
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	} else {
		t.co = co
	}
	defer func(co *client.ClientConn) {
		err := co.Close()
		if err != nil {
			log.Fatalf("Error close CoAP connection: %v", err)
		}
	}(t.co)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.ctx = ctx
	defer cancel()

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
