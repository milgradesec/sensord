package bluetooth

import (
	"encoding/binary"
	"strconv"

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
	var (
		value     string
		lastValue string
	)
	for {
		value = <-ds.ch
		if value != lastValue {
			lastValue = value
			n, _ := strconv.ParseUint(value, 10, 16)
			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(n))

			// fmt.Printf("Value: %s, Bytes -> %v\n", value, b)
			ds.char.Write(b) //nolint
		}
	}
}
