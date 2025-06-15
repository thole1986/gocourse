package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Person struct
type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// Sample data
var personData = map[string]Person{
	"1": {Name: "John Doe", Age: 30},
	"2": {Name: "Jane Doe", Age: 28},
	"3": {Name: "Jack Doe", Age: 25},
}

// Handler function for the endpoint
func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	// Check if the ID exists in the personData map
	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the person data to JSON and write to the response
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// Main function
func main() {
	// Define the port number in a variable
	port := 8000

	// Load the TLS certificate and key
	certFile := "cert.pem"
	keyFile := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom HTTP server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		TLSConfig: tlsConfig,
		// TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	// Set up the endpoint and the handler function
	http.HandleFunc("/person", getPersonHandler)

	// Enable HTTP/2 support
	// http2.ConfigureServer(server, &http2.Server{})

	// Start the server and print the confirmation message
	fmt.Printf("Server is running on port %d\n", port)
	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
