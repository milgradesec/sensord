package serial

import (
	"strings"

	"go.bug.st/serial"
)

var (
	defaultPort = "/dev/ttyACM0"
)

type Reader struct {
	port serial.Port
	ch   chan string
}

func NewReader(ch chan string) (*Reader, error) {
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open(defaultPort, mode)
	if err != nil {
		return nil, err
	}

	return &Reader{
		port: port,
		ch:   ch,
	}, nil
}

func (r *Reader) Start() {
	buff := make([]byte, 2048)
	for {
		n, err := r.port.Read(buff)
		if err != nil {
			break
		}

		if n == 0 {
			r.ch <- "EOF"
			break
		}

		line := string(buff[:n])
		if line != "\r\n" {
			r.ch <- strings.TrimSuffix(line, "\r\n")
		}
	}
}

func (r *Reader) Close() error {
	return r.port.Close()
}
