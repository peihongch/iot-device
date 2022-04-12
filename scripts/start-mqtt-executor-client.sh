#!/bin/bash

# start-mqtt-client.sh

device=$1
mode=$2
remote=$3
token=$4
topic=$5

./iot-device "$device" \
  --mode "$mode" \
  --name "$topic" \
  --remote "$remote" \
  --token "$token" &
