package ble

import (
	"github.com/rs/zerolog/log"
	"tinygo.org/x/bluetooth"
)

var (
	adapter    = bluetooth.DefaultAdapter
	deviceName = "AgroSensor"
)

func EnableAdapter() error {
	if err := adapter.Enable(); err != nil {
		return err
	}
	return nil
}

func StartGATTService() error {
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

	var heartRateCharacteristic bluetooth.Characteristic
	hrService := &HeartRateService{
		heartRate: &heartRateCharacteristic,
	}
	if err = adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.ServiceUUIDHeartRate,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &heartRateCharacteristic,
				UUID:   bluetooth.CharacteristicUUIDHeartRateMeasurement,
				Value:  []byte{0, heartRate},
				Flags:  bluetooth.CharacteristicNotifyPermission,
			},
		},
	}); err != nil {
		return err
	}
	log.Info().Msg("HearRate service running...")

	hrService.Handler()
	return nil
}
