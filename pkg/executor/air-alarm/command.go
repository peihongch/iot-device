package air_alarm

import (
	"fmt"
	"github.com/peihongch/iot-device/pkg"
)

type AirAlarmCommand struct {
	Timestamp int     `json:"ts"`
	Device    string  `json:"device"`
	Co        float64 `json:"co"`
	Threshold float64 `json:"threshold"`
}

func ExecCmd(cmd *AirAlarmCommand) {
	if cmd.Co >= cmd.Threshold {
		fmt.Println(pkg.SprintRed(fmt.Sprintf("!!!WARNING!!! 一氧化碳浓度到达 %v，危险阈值为 %v!", cmd.Co, cmd.Threshold)))
	}
}
