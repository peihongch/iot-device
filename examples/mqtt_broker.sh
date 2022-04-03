#!/bin/sh

./bin/iot-device air-contitioner --data ./data/devices/00:0f:00:70:91:0a/iot_telemetry_humidity.csv --name "v1/devices/me/telemetry" --remote tcp://172.19.241.103:11883 --token DHT22
