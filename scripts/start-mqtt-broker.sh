#!/bin/sh

device=$1
port=$2
topic=$3

./bin/iot-device "$device" --port "$port" --topic "$topic" >>log/"$device".log &
