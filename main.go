package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
