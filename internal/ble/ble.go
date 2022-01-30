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

func StartGATTService(ch chan string) error {
	adv := adapter.DefaultAdvertisement()

	serviceUUID, err := bluetooth.ParseUUID(serviceUUID)
	if err != nil {
		return err
	}

	err = adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: deviceName,
		ServiceUUIDs: []bluetooth.UUID{
			serviceUUID,
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
	charUUID, err := bluetooth.ParseUUID(charUUID)
	if err != nil {
		return err
	}

	if err = adapter.AddService(&bluetooth.Service{
		UUID: serviceUUID,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &distanceCharacteristic,
				UUID:   charUUID,
				Value:  []byte{0, 0},
				Flags:  bluetooth.CharacteristicNotifyPermission | bluetooth.CharacteristicReadPermission,
			},
		},
	}); err != nil {
		return err
	}

	log.Info().Msg("Distance service running...")

	for {
		distanceCharacteristic.Write([]byte(<-ch)) //nolint
	}
}

var (
	serviceUUID = "9298dcb2-47b1-4cb5-8dfa-3c865ea8163e"
	charUUID    = "5bb31353-d8cd-4d18-a22a-e735e23b5bdc"
)
