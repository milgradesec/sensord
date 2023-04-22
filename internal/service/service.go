package service

import (
	"github.com/kardianos/service"
	"github.com/milgradesec/sensord/internal/bluetooth"
	"github.com/milgradesec/sensord/internal/serial"
	"github.com/rs/zerolog/log"
)

type Service struct {
	Development bool
}

// Start implements the service.Service interface.
func (s *Service) Start(svc service.Service) error {
	go func() {
		s.run()
	}()
	return nil
}

func (s *Service) run() {
	ch := make(chan string)

	if s.Development {
		log.Info().Msg("Development mode enabled, using mock serial reader.")

		sr := serial.NewMockReader(ch)
		go sr.Start()
	} else {
		sr, err := serial.NewReader(ch)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create serial reader")
		}
		go sr.Start()
	}

	if err := bluetooth.StartGATTService(ch); err != nil {
		log.Fatal().Err(err).Msg("Failed to start GATT service")
	}
}

// Stop implements the service.Service interface.
func (s *Service) Stop(svc service.Service) error {
	return nil
}
