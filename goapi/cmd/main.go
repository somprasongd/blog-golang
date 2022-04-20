package main

import (
	"log"
	"net/http"
)

func main() {
	// starting server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
