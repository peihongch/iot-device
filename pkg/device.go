package pkg

// Device 传感器设备接口
type Device interface {
	// Start 启动设备守护进程
	Start()
	// Collect 传感器设备采集环境数据
	Collect()
}

// NewThermometer 实例化温度计设备
//  source 数据源
//  remote 数据发送的远端目的平台
func NewThermometer(source string, remote string) *Thermometer {
	return &Thermometer{}
}

// Thermometer 温度计
type Thermometer struct {
}

func (t Thermometer) Collect() {
	//TODO implement me
	panic("implement me")
}

func (t Thermometer) Start() {
	//TODO implement me
	panic("implement me")
}
