package main

import (
	"log"
	"net/http"
)

// create a file server and run it
func main() {
	// start the server
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./static"))))
}
