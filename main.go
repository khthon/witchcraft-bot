package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

var lineBotConfig LineBotConfig
var lineBotClient *linebot.Client
var err error

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Port: " + port)

	lineBotConfig = GetLineBotConfig()
	lineBotClient, err = linebot.New(lineBotConfig.ChannelSecret, lineBotConfig.ChannelAccessToken)
	if err != nil {
		log.Fatal("Can't create line bot")
	}

	http.HandleFunc("/linewebhook", lineWebHook)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func lineWebHook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		events, err := lineBotClient.ParseRequest(r)
		if err != nil {
			http.Error(w, "Invalid request method.", 405)
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				log.Println(event.ReplyToken)
			}
		}
	default:
		http.Error(w, "Invalid request method.", 405)
	}
}

func validateSignature(channelSecret, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))

	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}