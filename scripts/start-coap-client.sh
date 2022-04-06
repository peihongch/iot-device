#!/bin/bash

device=$1
device_id=$2
host=$3
port=$4
path=$5

./iot-device "$device" \
  --data ./devices/"$device_id"/iot_telemetry_"$device".csv \
  --host "$host" --port "$port" --path "$path" &
