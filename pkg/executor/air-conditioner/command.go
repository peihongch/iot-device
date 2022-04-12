package air_conditioner

import (
	"fmt"
	"github.com/fatih/color"
)

type AirConditionerCommand struct {
	Method string `json:"method"`
	Params struct {
		Device    string `json:"device"`
		Timestamp string `json:"ts"`
		Mode      string `json:"mode"`
		Target    int    `json:"target"`
	} `json:"params"`
}

func ExecCmd(cmd *AirConditionerCommand) {
	switch cmd.Params.Mode {
	case "hot":
		fmt.Printf("空调【模式：%s】【温度：%v℃】\n", color.RedString("制热"), color.CyanString(fmt.Sprintf("%v", cmd.Params.Target)))
	case "cold":
		fmt.Printf("空调【模式：%s】【温度：%v℃】\n", color.BlueString("制冷"), color.CyanString(fmt.Sprintf("%v", cmd.Params.Target)))
	case "dry":
		fmt.Printf("空调【模式：%s】\n", color.GreenString("除湿"))
	default:
		//bytes, err := json.Marshal(cmd)
		//if err != nil {
		//	return
		//}
		//color.Yellow("air-conditioner for test, %v", string(bytes))
	}
}
