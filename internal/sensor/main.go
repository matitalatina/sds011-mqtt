package sensor

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ryszard/sds011/go/sds011"
)

type Config struct {
	Topic          string
	SensorPortPath string
	CycleMinutes   uint8
	MqttBroker     string
}

func Start(c Config) {
	opts := mqtt.NewClientOptions().AddBroker(c.MqttBroker)
	opts.AutoReconnect = true
	opts.SetKeepAlive(30 * time.Second)
	opts.SetPingTimeout(10 * time.Second)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// sensor, err := sds011.New(c.SensorPortPath)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer sensor.Close()

	// if err = sensor.SetCycle(c.CycleMinutes); err != nil {
	// 	log.Printf("ERROR: sensor.SetCycle: %v", err)
	// }

	for {
		// point, err := sensor.Get()
		var noError error

		point, err := sds011.Point{
			PM10:      float64(rand.Intn(20)),
			PM25:      float64(rand.Intn(20)),
			Timestamp: time.Now(),
		}, noError
		time.Sleep(5 * time.Second)

		if err != nil {
			log.Printf("ERROR: sensor.Get: %v", err)
			continue
		}

		fmt.Fprintf(os.Stdout, "%v,%v,%v\n", point.Timestamp.Format(time.RFC3339), point.PM25, point.PM10)

		pointJSON, err := json.Marshal(point)

		if err != nil {
			log.Printf("ERROR: Marshal: %v", err)
			continue
		}

		if token := client.Publish(c.Topic, 0, false, pointJSON); token.Wait() && token.Error() != nil {
			fmt.Print(token.Error())
		}
	}
}
