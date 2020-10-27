package main

import "mattianatali.it/sds011-mqtt/internal/sensor"

func main() {
	c := sensor.Config{
		Topic:          "home/serina/edge-rpi/air-quality",
		SensorPortPath: "/dev/ttyUSB0",
		CycleMinutes:   1,
		MqttBroker:     "tcp://192.168.1.117:1883",
	}
	sensor.Start(c)
}
