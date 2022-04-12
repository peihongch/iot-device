package air_alarm

import (
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"strconv"
)

type AirAlarmCommand struct {
	Method string `json:"method"`
	Params struct {
		Device    string `json:"device"`
		Timestamp string `json:"ts"`
		Co        string `json:"co"`
		Threshold int    `json:"threshold"`
	} `json:"params"`
}

func ExecCmd(cmd *AirAlarmCommand) {
	co, err := strconv.ParseFloat(cmd.Params.Co, 64)
	if err != nil {
		log.Fatalln("invalid format of CO value", cmd.Params.Co)
	}

	if co >= float64(cmd.Params.Threshold) {
		color.Red("!!!WARNING!!! 一氧化碳浓度到达 %v，危险阈值为 %v!", cmd.Params.Co, cmd.Params.Threshold)
	} else {
		bytes, err := json.Marshal(cmd)
		if err != nil {
			return
		}
		color.Yellow("air-alarm for test, %v", string(bytes))
	}
}
