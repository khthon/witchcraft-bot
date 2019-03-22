package main

import (
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

var lineBotConfig LineBotConfig
var lineBotClient *linebot.Client
var err error

type WebhookTextMessage struct {
	replyToken string `json:"replyToken"`
	messageType string `json:"type"`
	//message TextMessage `json:"message"`
}

//type TextMessage struct {
//	id string `json:"id"'`
//	messageType string `json:"type"'`
//	text string `json:"text"'`
//}

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
		defer r.Body.Close()
		var webhookTextMessage WebhookTextMessage
		err := json.NewDecoder(r.Body).Decode(&webhookTextMessage)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//body, err := ioutil.ReadAll(r.Body)

		//if err != nil {
		//	http.Error(w, err.Error(), 500)
		//	return
		//}
		//
		//log.Println(body)
		//
		//var webhookTextMessage WebhookTextMessage
		//err = json.Unmarshal(body, &webhookTextMessage)
		//if err != nil {
		//	http.Error(w, err.Error(), 500)
		//	return
		//}

		log.Println(webhookTextMessage)

		if _, err := lineBotClient.ReplyMessage(webhookTextMessage.replyToken, linebot.NewTextMessage(webhookTextMessage.message.text)).Do(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	default:
		w.WriteHeader(405)
		http.Error(w, "Invalid request method.", 405)
	}


}