all: data main.go
	@go build main.go -o iot-device

.PHONY: data
data: data/iot_telemetry_data.csv data/preprocessing.py
	@cd data && python3 preprocessing.py && cd -

.PHONY: clean
clean:
	@rm -rf iot-device data/devices