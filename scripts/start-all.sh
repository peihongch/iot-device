#!/bin/bash

mkdir -p log

device_id=$1
./start-mqtt-client.sh "th" "$device_id" "tcp://172.19.241.103:11883" "temp_humidity_sensor_$device_id" "v1/devices/me/telemetry"
./start-mqtt-client.sh "gas" "$device_id" "tcp://172.19.241.103:11883" "gas_sensor_$device_id" "v1/devices/me/telemetry"
./start-mqtt-client.sh "light" "$device_id" "tcp://172.19.241.103:11883" "light_sensor_$device_id" "v1/devices/me/telemetry"
echo "MQTT sensor devices started!"

./start-coap-client.sh "motion" "$device_id" "172.19.241.103" "5684" "api/v1/DHT22/telemetry"
echo "COAP sensor device started!"

port=1883
for device in air-conditioner air-alarm; do
  ./start-mqtt-broker.sh "$device" "$port" "v1/execution"
  port=$(expr $port + 1)
done
echo "MQTT executor devices started!"
