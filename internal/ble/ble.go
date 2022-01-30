package ble

import (
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
	"tinygo.org/x/bluetooth"
)

var (
	adapter          = bluetooth.DefaultAdapter
	deviceName       = "AgroSensor"
	heartRate  uint8 = 75
)

func EnableAdapter() error {
	if err := adapter.Enable(); err != nil {
		return err
	}
	return nil
}

func StartGATTService() error {
	var (
		heartRateMeasurement bluetooth.Characteristic
	)

	adv := adapter.DefaultAdvertisement()
	err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    deviceName,
		ServiceUUIDs: []bluetooth.UUID{bluetooth.ServiceUUIDHeartRate},
	})
	if err != nil {
		return err
	}

	if err := adv.Start(); err != nil {
		return err
	}
	log.Info().Msgf("Advertising device as '%s'", deviceName)

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
		return err
	}

	nextBeat := time.Now()
	for {
		nextBeat = nextBeat.Add(time.Minute / time.Duration(heartRate))
		time.Sleep(time.Until(nextBeat))

		heartRate = randomInt(65, 85)
		heartRateMeasurement.Write([]byte{0, heartRate}) //nolint
	}
}

func randomInt(min, max int) uint8 {
	return uint8(min + rand.Intn(max-min)) //nolint
}
