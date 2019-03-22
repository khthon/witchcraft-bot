package main

import (
	"fmt"
	"html"
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

	http.HandleFunc("/linewebhook", lineWebHook)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func lineWebHook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", 405)
	}

	fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))
}