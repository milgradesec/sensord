package ble

import (
	"math/rand"
	"time"

	"tinygo.org/x/bluetooth"
)

var (
	heartRate uint8 = 75
)

type HeartRateService struct {
	heartRate *bluetooth.Characteristic
}

func (hrs *HeartRateService) Handler() {
	nextBeat := time.Now()
	for {
		nextBeat = nextBeat.Add(time.Minute / time.Duration(heartRate))
		time.Sleep(time.Until(nextBeat))

		heartRate = randomInt(65, 85)
		hrs.heartRate.Write([]byte{0, heartRate}) //nolint
	}
}

func randomInt(min, max int) uint8 {
	return uint8(min + rand.Intn(max-min)) //nolint
}
