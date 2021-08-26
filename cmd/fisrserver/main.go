package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mohakamg/fastinversesquareroot/fisr"
)

// Declare CLI Flags
var (
	// Only one of the flags can be specified, if both are specified the server will fail to start

	// The number whose inverse square root has to be calculated
	number = flag.String("number", "", "Number to calculate the Inverse Square root Of")
	// The endpoint location where the server needs to start
	serverEndpoint = flag.String("server-endpoint", "localhost:30001", "The endpoint of the server")
)

// Declare a struct to describe the request
type FISRRequest struct {
	// The values field describes a list
	// of 32 bit floating point values whose
	// Fast Inverse Square Root needs to be calculated.
	Instances []float32 `json:"instances"`
}

// Declare a struct to describe the response
type FISRResponse struct {
	// A map that contains the set of the values in the request
	// mapped to its inverse square root
	Result map[string]float32 `json:"result"`
}

// A preflight function before the main thread
// is executed.
func preFlight() {
	// Parse the flags
	flag.Parse()
}

func main() {

	// Do the preflight Setup
	preFlight()

	// Check if this is a standalone mode
	if *number != "" {
		value, err := strconv.ParseFloat(*number, 32)
		if err != nil {
			log.Fatalf("Error while parsing input: ", err)
		}
		if inverseSquareRoot, err := fisr.FastInverseSquareRoot(float32(value)); err != nil {
			log.Fatalf("Error while calculating: ", err)
		} else {
			log.Println("Result: ", inverseSquareRoot)
		}
	} else {
		// Declare the routes
		http.HandleFunc("/fisr", func(writer http.ResponseWriter, req *http.Request) {
			// This function is used to compute the inverse square roots
			// of the incoming request
			switch req.Method {
			case "POST":
				// Deserialize the Request
				log.Println("Incoming Request: ", req.Body)

				// Set the Response Content type
				writer.Header().Set("Content-Type", "application/json")

				// Create a placeholder object in memory
				var incomingReq FISRRequest

				// Desealizing using the decoder
				err := json.NewDecoder(req.Body).Decode(&incomingReq)
				log.Println("Decoded Request: ", incomingReq)

				// If error in decoding incoming request
				if err != nil {
					writer.WriteHeader(500)
					writer.Write([]byte("Could not decode incoming request"))
					log.Println("Could not decode incoming request: ", err)
					return
				}

				// Compute the response
				var responseObj FISRResponse
				responseObj.Result = make(map[string]float32)
				for _, instance := range incomingReq.Instances {
					inverseSquareRoot, err := fisr.FastInverseSquareRoot(float32(instance))
					if err != nil {
						writer.WriteHeader(500)
						writer.Write([]byte("Could not compute fisr"))
						log.Println("Could not compute fisr: ", err)
						return
					}
					responseObj.Result[fmt.Sprint(instance)] = inverseSquareRoot
				}
				serializedResp, err := json.Marshal(responseObj)
				if err != nil {
					writer.WriteHeader(500)
					writer.Write([]byte("Could not Serialize Response"))
					log.Println("Could not Serialize Response: ", err)
				}
				writer.Write(serializedResp)

			default:
				log.Println("%s Method not supported by server", req.Method)
			}
		})

		// Start the Server
		log.Println("Starting server: ", *serverEndpoint)
		http.ListenAndServe(*serverEndpoint, nil)

	}

}
