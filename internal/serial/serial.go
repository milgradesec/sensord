package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func readDataFromSerial() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
	} else {
		for _, port := range ports {
			fmt.Printf("Found port: %v\n", port)
		}
	}

	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open("/dev/ttyACM0", mode)
	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))
	}
}
