package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var lineBot LineBot

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Port: " + port)

	http.HandleFunc("/linewebhook", handler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "Hello, World!")
}