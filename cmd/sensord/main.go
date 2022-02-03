package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/milgradesec/sensord/internal/bluetooth"
	"github.com/milgradesec/sensord/internal/serial"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	debug := flag.Bool("debug", false, "Enable debug logging.")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Info().Msgf("Sensord %s", Version)
	log.Info().Msgf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())

	ch := make(chan string)
	sr, err := serial.NewReader(ch)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create serial reader")
	}
	go sr.Start()

	if err := bluetooth.StartGATTService(ch); err != nil {
		log.Fatal().Err(err).Msg("Failed to start GATT service")
	}
}

var (
	Version = "DEV"
)
