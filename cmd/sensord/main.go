package main

import (
	"os"
	"runtime"

	"github.com/milgradesec/sensord/internal/bluetooth"
	"github.com/milgradesec/sensord/internal/serial"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Msgf("Sensord %s", Version)
	log.Info().Msgf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())

	ch := make(chan string)
	sr, err := serial.NewReader(ch)
	if err != nil {
		log.Fatal().Msgf("failed to create serial reader: %v", err)
	}
	go sr.Start()

	if err := bluetooth.StartGATTService(ch); err != nil {
		log.Fatal().Msgf("failed to start GATT service: %v", err)
	}
}

var (
	Version = "DEV"
)
