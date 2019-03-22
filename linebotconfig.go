package main

import (
	"log"
	"os"
)

type LineBotConfig struct {
	ChannelID string
	ChannelSecret string
	ChannelAccessToken string
}

func GetLineBotConfig() LineBotConfig {
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

	return LineBotConfig{
		ChannelID: lineChannelID,
		ChannelSecret:lineChannelSecret,
		ChannelAccessToken: lineChannelAccessToken,
	}
}

