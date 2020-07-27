package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
)

func PostDMToMe(text string) error {
	text = EscapeProcessing(text)
	url := "https://api.twitter.com/1.1/direct_messages/events/new.json"
	payload := strings.NewReader("{\n\t\"event\":{\n\t\t\t\"type\":\"message_create\",\n\t\t\"message_create\": {\n    \"target\": {\n      \"recipient_id\": \"1284149601826598913\"\n    },\n    \"message_data\": {\n      \"text\": \"" + text + "\"\n    }\n  }\n\t}\n}")
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	req.Header.Add("content-type", "application/json")

	if _, err := httpClient.Do(req); err != nil {
		return err
	}

	return nil
}
