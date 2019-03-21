package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println(port)

	lineChannelID := os.Getenv("LineChannelID")
	if lineChannelID == "" {
		log.Fatal("$LineChannelID must be set")
	}

	log.Println(lineChannelID)

	lineChannelAccessToken := os.Getenv("LineChannelAccessToken")
	if lineChannelAccessToken == "" {
		log.Fatal("LineChannelAccessToken must be set")
	}

	log.Println(lineChannelAccessToken)

	lineChannelSecret := os.Getenv("LineChannelSecret")
	if lineChannelSecret == "" {
		log.Fatal("LineChannelSecret must be set")
	}

	log.Println(lineChannelSecret)
}