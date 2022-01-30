package ble

import (
	"github.com/rs/zerolog/log"
	"tinygo.org/x/bluetooth"
)

type GATTServer struct {
	adv *bluetooth.Advertisement
}

func (gs *GATTServer) Start() error {
	err := gs.adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: deviceName,
		ServiceUUIDs: []bluetooth.UUID{
			bluetooth.NewUUID(serviceUUID),
		},
	})
	if err != nil {
		return err
	}

	if err := gs.adv.Start(); err != nil {
		return err
	}
	log.Info().Msgf("Advertising device as '%s'", deviceName)

	return nil
}
