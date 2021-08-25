package main

import (
	"flag"

	"github.com/mohakamg/fastinversesquareroot/fisr"
)

// Declare CLI Flags
var (
	// Only one of the flags can be specified, if both are specified the server will fail to start

	// The number whose inverse square root has to be calculated
	number = flag.Float32("number", nil, "Number to calculate the Inverse Square root Of")
	// The endpoint location where the server needs to start
	serverEndpoint = flag.String("server-endpoint", "localhost:30001", "The endpoint of the server")
)

// A preflight function before the main thread
// is executed.
func preFlight() {
	flag.Parse()
}

func main() {

	preFlight()

	if number != nil {
		if inverseSquareRoot, err := fisr.FastInverseSquareRoot(number); err != nil {
			log.FatalF("Error while calculating: ", err)
		} else {
			log.Println("Result: ", inverseSquareRoot)
		}
	}

}
