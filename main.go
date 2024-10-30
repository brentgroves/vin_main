package main

import (
	"log"
	// greetings "github.com/brentgroves/greetings3" // should get v0.3.0
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("vin_main: ")
	log.SetFlags(0)

	// Request greeting messages for the names.
	// messages, err := greetings.Hellos(names)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// If no error was returned, print the returned map of
	// messages to the console.
	// fmt.Println(messages)
}
