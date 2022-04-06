#!/bin/sh

mkdir -p log

device_id=$1
for device in temp-humitidy gas light; do
  ./start-mqtt-client.sh "$device" "$device_id" "tcp://172.19.241.103:11883" "DHT22" "v1/devices/me/telemetry"
done
echo "MQTT sensor devices started!"

./start-coap-client.sh "motion" "$device_id" "172.19.241.103" "5684" "api/v1/DHT22/telemetry"
echo "COAP sensor device started!"

port=1883
for device in air-conditioner air-alarm; do
  ./start-mqtt-broker.sh "$device" "$port" "v1/execution"
  port=$(expr $port + 1)
done
echo "MQTT executor devices started!"
