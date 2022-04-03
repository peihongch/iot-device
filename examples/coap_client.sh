#!/bin/sh

./bin/iot-device motion --data ./data/devices/00:0f:00:70:91:0a/iot_telemetry_motion.csv --path "api/v1/DHT22/telemetry" --host 172.19.241.103 --port 5684
