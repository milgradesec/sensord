//go:build linux

package bluetooth

import (
	"encoding/binary"
	"strconv"

	"github.com/rs/zerolog/log"
	"tinygo.org/x/bluetooth"
)

var (
	distanceUUID = [16]byte{0x5b, 0xb3, 0x13, 0x53, 0xd8, 0xcd, 0x4d, 0x18, 0xa2, 0x2a, 0xe7, 0x35, 0xe2, 0x3b, 0x5b, 0xdc}
)

type DistanceService struct {
	char *bluetooth.Characteristic
	ch   chan string
}

func (ds *DistanceService) Handler() {
	log.Info().Str("service", ds.Name()).Msg("Service running")

	var (
		value string
	)
	for {
		value = <-ds.ch

		n, _ := strconv.ParseUint(value, 10, 16)
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(n))
		ds.char.Write(b) //nolint
	}
}

func (ds *DistanceService) Name() string {
	return "distance"
}
