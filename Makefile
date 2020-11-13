.PHONY: build build-arm

BUILD_CMD=go build -o dist/pm-sensor cmd/sensor/main.go

build-arm6:
	GOOS=linux GOARCH=arm GOARM=6 $(BUILD_CMD)

build-arm:
	GOOS=linux GOARCH=arm GOARM=7 $(BUILD_CMD)

build:
	$(BUILD_CMD)

cp-edge:
	scp dist/pm-sensor edge-rpi:pm-sensor
	