package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// /
func handler(response http.ResponseWriter, request *http.Request) {
	// Log the request
	log.Printf("Received %s request for %s from %s", request.Method, request.URL, request.Host)

	// Create response string builder and add some metadata
	var responseHeaderListBuilder strings.Builder
	responseHeaderListBuilder.WriteString(fmt.Sprintf("Your IP: %s\n%s %s\n\n", request.Host, request.Method, request.URL))

	// Iterate over each key-value pair (eg "Content-Type", "Accept")
	for reqHeaderName, reqHeaders := range request.Header {

		// Iterate over each value in reqHeaders[] and write to builder
		for _, reqHeaderValue := range reqHeaders {
			// "HeaderName: HeaderValue\n"
			responseHeaderListBuilder.WriteString(fmt.Sprintf("%v: %v\n", reqHeaderName, reqHeaderValue))
		}
	}

	// Convert builder to string and write to response
	var responseHeaderList string = responseHeaderListBuilder.String()
	fmt.Fprint(response, responseHeaderList)
}

// Main
func main() {
	// The goal is to spin up a simple server for testing web apps
	var port string = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Log
	log.Printf("Serving website on port %s\n", port)

	// Serve Website
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
