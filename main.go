package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

var heartRate uint8 = 75 // 75bpm

func main() {
	fmt.Println("starting")

	if err := adapter.Enable(); err != nil {
		panic(err)
	}

	adv := adapter.DefaultAdvertisement()
	err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    "AGROSENSOR",
		ServiceUUIDs: []bluetooth.UUID{bluetooth.ServiceUUIDHeartRate},
	})
	if err != nil {
		panic(err)
	}

	if err := adv.Start(); err != nil {
		panic(err)
	}

	var heartRateMeasurement bluetooth.Characteristic
	err = adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.ServiceUUIDHeartRate,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &heartRateMeasurement,
				UUID:   bluetooth.CharacteristicUUIDHeartRateMeasurement,
				Value:  []byte{0, heartRate},
				Flags:  bluetooth.CharacteristicNotifyPermission,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	nextBeat := time.Now()
	for {
		fmt.Println("Time:" + time.Now().Format("04:05.000") + " -- Value: " + strconv.FormatUint(uint64(heartRate), 10))

		nextBeat = nextBeat.Add(time.Minute / time.Duration(heartRate))
		time.Sleep(time.Until(nextBeat))

		heartRate = randomInt(65, 85)
		heartRateMeasurement.Write([]byte{0, heartRate}) //nolint
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) uint8 {
	return uint8(min + rand.Intn(max-min)) //nolint
}
