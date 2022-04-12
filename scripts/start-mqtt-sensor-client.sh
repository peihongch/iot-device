#!/bin/bash

# start-mqtt-client.sh

device=$1
device_id=$2
remote=$3
token=$4
topic=$5

./iot-device "$device" \
  --data ./devices/"$device_id"/iot_telemetry_"$device".csv \
  --name "$topic" \
  --remote "$remote" \
  --token "$token" &
