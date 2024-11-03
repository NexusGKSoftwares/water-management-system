package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Import a popular router package
)

func main() {
	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Water Management System!")
	})

	// Start the server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
