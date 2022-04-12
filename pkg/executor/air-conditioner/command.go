package air_conditioner

import (
	"fmt"
	"github.com/peihongch/iot-device/pkg"
)

type AirConditionerCommand struct {
	Mode   string `json:"mode"`
	Target int    `json:"target"`
}

func ExecCmd(cmd *AirConditionerCommand) {
	switch cmd.Mode {
	case "hot":
		fmt.Printf("空调【模式：%s】【温度：%v℃】\n", pkg.SprintRed("制热"), pkg.SprintCyan(fmt.Sprintf("%v", cmd.Target)))
	case "cold":
		fmt.Printf("空调【模式：%s】【温度：%v℃】\n", pkg.SprintBlue("制冷"), pkg.SprintCyan(fmt.Sprintf("%v", cmd.Target)))
	case "dry":
		fmt.Printf("空调【模式：%s】\n", pkg.SprintGreen("除湿"))
	}
}
