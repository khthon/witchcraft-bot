package main

import (
	"log"
	"os"
)

type LineBot struct {
	ChannelID string
	ChannelSecret string
	ChannelAccessToken string
}

func NewLineBot() LineBot {
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

	return LineBot{
		ChannelID: lineChannelID,
		ChannelSecret:lineChannelSecret,
		ChannelAccessToken: lineChannelAccessToken,
	}
}

func (lineBot LineBot) ReplyText(message string) {

}

