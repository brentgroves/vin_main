package main

import (
	"fmt"
	"log"

	vin "github.com/brentgroves/vin1/vin-stages/4"
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
		euSmallVIN = "W09000051T2123456"
	)

	// slice of vin or vinEU types that implement VIN interface
	var testVINs []vin.VIN
	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	// now there is no need to cast!
	testVINs = append(testVINs, testVIN)

	for _, vin := range testVINs {
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			log.Printf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
		}
		fmt.Println(manufacturer)

	}

}
