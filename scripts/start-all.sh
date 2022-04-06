#!/bin/bash

mkdir -p log

declare -a device_id_map=(["1"]="00:0f:00:70:91:0a" ["2"]="1c:bf:ce:15:ec:4d" ["3"]="b8:27:eb:bf:9d:51")
declare -a device_token_map=(["th"]="temp_humidity_sensor" ["gas"]="gas_sensor" ["light"]="light_sensor")

device_id=$1
for device in th gas light; do
  ./start-mqtt-client.sh "$device" "${device_id_map[$device_id]}" "tcp://172.19.241.103:11883" "${device_token_map[$device]}_$device_id" "v1/devices/me/telemetry"
done
echo "MQTT sensor devices started!"

./start-coap-client.sh "motion" "${device_id_map[$device_id]}" "172.19.241.103" "5684" "api/v1/DHT22/telemetry"
echo "COAP sensor device started!"

port=1883
for device in air-conditioner air-alarm; do
  ./start-mqtt-broker.sh "$device" "$port" "v1/execution"
  port=$(expr $port + 1)
done
echo "MQTT executor devices started!"
