package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	// mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Root handler")
	// })

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Users handler")
	})
	mux.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Items handler")
	})

	// "/app/users"
	// "/app/items"

	app := http.NewServeMux()
	app.Handle("/app/", http.StripPrefix("/app", mux))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Server started on port 8080")
	server.ListenAndServe()

}
