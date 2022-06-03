package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/kardianos/service"

	sensord "github.com/milgradesec/sensord/internal/service"
)

func main() { //nolint
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	var (
		debug       = flag.Bool("debug", false, "Enable debug logging.")
		dev         = flag.Bool("dev", false, "Enable development mode.")
		serviceFlag = flag.String("service", "", "Manage system service.")
		version     = flag.Bool("version", false, "Show version information.")
	)
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if *version {
		log.Info().Msgf("Sensord %s, %s, %s", Version, Commit, BuildTime)
		log.Info().Msgf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
		return
	}

	log.Info().Msgf("Sensord %s", Version)
	log.Info().Msgf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())

	svcConfig := &service.Config{
		Name:        "sensord",
		DisplayName: "sensord",
		Description: "Sensord service",

		Dependencies: []string{
			"After=bluetooth.target",
			"Requires=bluetooth.target",
		},
	}

	s := &sensord.Service{
		Development: *dev,
	}
	svc, err := service.New(s, svcConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create service")
	}

	if *serviceFlag != "" {
		if err := service.Control(svc, *serviceFlag); err != nil {
			log.Fatal().Err(err).Msg("service control error")
		}

		switch *serviceFlag {
		case "install":
			log.Info().Msg("Service installed successfully.")
		case "uninstall":
			log.Info().Msg("Service removed successfully.")
		case "start":
			log.Info().Msg("Service started.")
		case "stop":
			log.Info().Msg("Service stopped.")
		case "restart":
			log.Info().Msg("Service restarted.")
		default:
			log.Error().Msgf("invalid argument: %s", *serviceFlag)
		}
		return
	}
	if err := svc.Run(); err != nil {
		log.Fatal().Err(err).Msg("failed to start service")
	}
}

var (
	Version   = "DEV"
	Commit    = ""
	BuildTime = ""
)
