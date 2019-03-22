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
	switch r.Method {
	case http.MethodPost:
		fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))
	default:
		w.WriteHeader(405)
		http.Error(w, "Invalid request method.", 405)
	}


}