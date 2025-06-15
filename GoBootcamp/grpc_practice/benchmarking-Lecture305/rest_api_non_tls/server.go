package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting greet handler")
	if r.Method != http.MethodPost {
		log.Println("Wrong method")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Bad request")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Validation
	if len(req.Name) < 5 || len(req.Name) > 50 || !regexp.MustCompile("^[a-zA-Z]+$").MatchString(req.Name) {
		log.Println(req.Name)
		log.Println(len(req.Name))
		log.Println(!regexp.MustCompile("^[a-zA-Z]+$").MatchString(req.Name))
		log.Println("Validation failed")
		http.Error(w, "Invalid request: name must be 5-50 characters long and contain only letters", http.StatusBadRequest)
		return
	}

	resp := HelloResponse{Message: "Hello, " + req.Name}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	log.Println("Sent response")
}

func main() {
	http.HandleFunc("/v1/greet", greetHandler)
	log.Println("HTTP server is running on port :8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
