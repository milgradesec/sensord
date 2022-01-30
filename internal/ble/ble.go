package ble

import (
	"github.com/rs/zerolog/log"
	"tinygo.org/x/bluetooth"
)

var (
	adapter     = bluetooth.DefaultAdapter
	deviceName  = "AgroSensor"
	serviceUUID = [16]byte{0x92, 0x98, 0xdc, 0xb2, 0x47, 0xb1, 0x4c, 0xb5, 0x8d, 0xfa, 0x3c, 0x86, 0x5e, 0xa8, 0x16, 0x3e}
)

func EnableAdapter() error {
	if err := adapter.Enable(); err != nil {
		return err
	}
	return nil
}

func StartGATTService(ch chan string) error {
	adv := adapter.DefaultAdvertisement()

	err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: deviceName,
		ServiceUUIDs: []bluetooth.UUID{
			bluetooth.NewUUID(serviceUUID),
		},
	})
	if err != nil {
		return err
	}

	if err := adv.Start(); err != nil {
		return err
	}
	log.Info().Msgf("Advertising device as '%s'", deviceName)

	var distanceCharacteristic bluetooth.Characteristic
	distanceService := &DistanceService{
		char: &distanceCharacteristic,
		ch:   ch,
	}

	if err = adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.NewUUID(serviceUUID),
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &distanceCharacteristic,
				UUID:   bluetooth.NewUUID(distanceUUID),
				Value:  []byte{0, 0},
				Flags:  bluetooth.CharacteristicNotifyPermission,
			},
		},
	}); err != nil {
		return err
	}
	distanceService.Handler()

	return nil
}
