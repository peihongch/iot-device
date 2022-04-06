# iot-device

## 项目介绍

云边端一体架构中端设备模拟器。

## 使用指南

```shell
(venv) peihongchen in ~/iot-device on main ● λ ./bin/iot-device 
Usage:
  iot-device [command]

Available Commands:
  air-alarm       [mqtt] Start up air alarm device
  air-conditioner [mqtt] Start up air conditioner device
  completion      Generate the autocompletion script for the specified shell
  gas             [mqtt] Start up gas (co, smoke and lpg) sensor device
  help            Help about any command
  light           [mqtt] Start up light sensor device
  motion          [coap] Start up motion sensor device
  th              [mqtt] Start up th (temperature and humidity) sensor device

Flags:
      --data string   device data file (e.g. ./data.csv)
  -h, --help          help for iot-device
      --name string   device name

Use "iot-device [command] --help" for more information about a command.
```
iot-device支持的命令如下：
- air-alarm: 启动气体警报器，采用mqtt协议;
- air-conditioner: 启动空调, 采用mqtt协议;
- gas: 启动气体传感器，包含co、smoke和lpg指标，采用mqtt协议;
- light: 启动光线传感器，采用mqtt协议;
- motion: 启动移动传感器，采用coap协议;
- th: 启动温湿度传感器，包含temp和humidity指标，采用mqtt协议。