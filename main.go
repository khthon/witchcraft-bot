package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

var lineBotConfig LineBotConfig
var bot *linebot.Client
var err error

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Port: " + port)

	lineBotConfig = GetLineBotConfig()
	bot, err = linebot.New(lineBotConfig.ChannelSecret, lineBotConfig.ChannelAccessToken)
	if err != nil {
		log.Fatal("Can't create line bot")
	}

	http.HandleFunc("/linewebhook", lineWebHook)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func lineWebHook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		events, err := bot.ParseRequest(r)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	default:
		http.Error(w, "Invalid request method.", 405)
	}
}