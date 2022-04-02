#!/bin/sh

./bin/iot-device humidity --data ./data/devices/iot_telemetry_humidity.csv --name "v1/devices/me/telemetry" --remote tcp://172.19.241.103:11883 --token DHT22
