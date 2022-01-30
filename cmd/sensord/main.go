package main

import (
	"os"
	"runtime"

	"github.com/milgradesec/sensord/internal/ble"
	"github.com/milgradesec/sensord/internal/serial"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Msgf("Sensord %s", Version)
	log.Info().Msgf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())

	if err := ble.EnableAdapter(); err != nil {
		log.Fatal().Msgf("failed to enable BLE adapter: %v", err)
	}
	log.Info().Msg("BLE stack enabled")

	ch := make(chan string)
	go func() {
		if err := serial.StartReader(ch); err != nil {
			log.Fatal().Msgf("error reading data from serial port: %v", err)
		}
	}()

	if err := ble.StartGATTService(); err != nil {
		log.Fatal().Msgf("failed to start GATT service: %v", err)
	}
}

var (
	Version = "DEV"
)
