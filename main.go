package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Current request count
type Counter struct {
	count     int
	mutex     sync.Mutex
	startTime time.Time
}

// Increment counter
func (c *Counter) increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

// Get current count
func (c *Counter) getCount() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

// Create response
func responseBuilder(request *http.Request, currentCount int, startTime time.Time) string {

	// Create empty response body
	var responseBodyBuilder strings.Builder

	responseBodyBuilder.WriteString(fmt.Sprintf("Your Host: %s\n", request.Host))
	responseBodyBuilder.WriteString(fmt.Sprintf("Uptime: %s - Current Requests: %d\n", formatDuration(time.Since(startTime)), currentCount))
	responseBodyBuilder.WriteString(fmt.Sprintf("%s %s %s\n\n", request.Method, request.Proto, request.URL))

	// Now we build the response body header list
	// Iterate over each key-value pair (eg "Content-Type", "Accept")
	for reqHeaderName, reqHeaders := range request.Header {
		// Iterate over each value in reqHeaders[] and write to builder
		for _, reqHeaderValue := range reqHeaders {
			responseBodyBuilder.WriteString(fmt.Sprintf("%v: %v\n", reqHeaderName, reqHeaderValue)) // "HeaderName: HeaderValue\n"
		}
	}

	// Convert the string builder to a string and return it
	var responseBody string = responseBodyBuilder.String()
	return responseBody
}

// Main
func main() {
	// Create counter
	var counter *Counter = &Counter{
		startTime: time.Now(),
	}

	log.Println("Server Online! StartTime: ", counter.startTime)

	// Handle /
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		counter.increment()
		fmt.Fprint(response, responseBuilder(request, counter.getCount(), counter.startTime))
	})

	// Bind to port
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formatDuration(d time.Duration) string {
	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d:%02d", days, hours, minutes, seconds)
}
