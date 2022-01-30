package serial

import (
	"go.bug.st/serial"
)

var (
	defaultPort = "/dev/ttyACM0"
)

func StartReader(ch chan string) error {
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open(defaultPort, mode)
	if err != nil {
		return err
	}
	defer port.Close()

	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			return err
		}
		if n == 0 {
			ch <- "EOF"
			break
		}
		ch <- string(buff[:n])
	}

	return nil
}
