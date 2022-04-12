all: data iot-device

iot-device: main.go
	@go build -o bin/iot-device

.PHONY: data
data: data/iot_telemetry_data.csv data/preprocessing.py
	@cd data && python3 preprocessing.py && cd -

.PHONY: clean
clean:
	@rm -rf iot-device data/devices

.PHONY: dist
dist: data iot-device
	@rm -rf dist && mkdir -p dist
	@cp -r data/devices dist
	@cp scripts/*.sh dist
	@cp bin/iot-device dist
	@chmod a+x dist/*.sh
	@tree dist
	@echo "Done!"
