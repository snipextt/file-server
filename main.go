package main

import (
	"log"
	"net/http"
)

// create a web server
func main() {
	// start the server
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
