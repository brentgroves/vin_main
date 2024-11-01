package main

import (
	"fmt"
	"log"

	vin "github.com/brentgroves/vin1/vin-stages/3"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("vin_main: ")
	log.SetFlags(0)

	const (
		validVIN   = "W0L000051T2123456"
		invalidVIN = "W0"
	)

	vin1, err := vin.NewVIN(validVIN)
	if err != nil {
		t.Errorf("creating NewVIN from %s returned an error: %s", validVIN, err.Error())
	}

	manufacturer := vin1.Manufacturer()
	if manufacturer != "W0L" {
		log.Fatal("W0L error")
	}
	fmt.Println(manufacturer)

	// Request greeting messages for the names.
	// messages, err := greetings.Hellos(names)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// If no error was returned, print the returned map of
	// messages to the console.
	// fmt.Println(messages)
}
