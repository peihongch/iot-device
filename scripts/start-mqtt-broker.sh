#!/bin/bash

# start-mqtt-broker.sh

device=$1
port=$2
topic=$3

./iot-device "$device" --port "$port" --topic "$topic" &
