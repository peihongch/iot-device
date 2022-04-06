#!/bin/sh

device=$1
device_id=$2
host=$3
port=$4
path=$5

./bin/iot-device "$device" \
  --data ./data/devices/"$device_id"/iot_telemetry_"$device".csv \
  --host "$host" --port "$port" --path "$path" \
  >>log/"$device"-"$device_id".log &
