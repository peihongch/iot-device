#!/bin/bash

# 设备环境（即不同的服务器），取值1，2和3
device_id=$1

# 启动MQTT设备
./start-mqtt-sensor-client.sh "th" "$device_id" "tcp://172.19.241.103:11883" \
            "temp_humidity_sensor_$device_id" "v1/devices/me/telemetry"
./start-mqtt-sensor-client.sh "gas" "$device_id" "tcp://172.19.241.103:11883" \
            "gas_sensor_$device_id" "v1/devices/me/telemetry"
./start-mqtt-sensor-client.sh "light" "$device_id" "tcp://172.19.241.103:11883" \
            "light_sensor_$device_id" "v1/devices/me/telemetry"

echo "All MQTT sensor devices started!"

./start-mqtt-executor-client.sh "air-alarm" "client" "tcp://172.19.241.103:11883" \
            "air_alarm_$device_id" "v1/devices/me/rpc/request/+"
./start-mqtt-executor-client.sh "air-conditioner" "client" "tcp://172.19.241.103:11883" \
            "air_conditioner_$device_id" "v1/devices/me/rpc/request/+"

echo "All MQTT executor devices started!"

# 启动COAP设备
./start-coap-client.sh "motion" "$device_id" "172.19.241.103" "5684" \
            "api/v1/motion_sensor_$device_id/telemetry"
echo "COAP sensor device started!"

#port=1883
#for device in air-conditioner air-alarm; do
#  ./start-mqtt-broker.sh "$device" "$port" "v1/execution"
#  port=$(expr $port + 1)
#done
#echo "MQTT executor devices started!"
