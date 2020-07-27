package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dghubble/oauth1"
)

func UpdateProfile(v url.Values) {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.twitter.com/1.1/account/update_profile.json"
	req, _ := http.NewRequest("POST", url+"?"+v.Encode(), nil)

	if _, err := httpClient.Do(req); err != nil {
		log.Fatal(err)
	}
}

func UpdateName(Name string) {
	v := url.Values{}
	v.Add("name", Name)
	UpdateProfile(v)
}

func UpdateDescription(Description string) {
	v := url.Values{}
	v.Add("description", Description)
	UpdateProfile(v)
}
