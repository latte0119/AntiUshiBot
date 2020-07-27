package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/dghubble/oauth1"
)

func UpdateProfile(v url.Values) error {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.twitter.com/1.1/account/update_profile.json"
	req, err := http.NewRequest("POST", url+"?"+v.Encode(), nil)
	if err != nil {
		return err
	}

	if _, err := httpClient.Do(req); err != nil {
		return err
	}
	return nil
}

func UpdateName(Name string) error {
	v := url.Values{}
	v.Add("name", Name)
	if err := UpdateProfile(v); err != nil {
		return err
	}
	return nil
}

func UpdateDescription(Description string) error {
	v := url.Values{}
	v.Add("description", Description)
	if err := UpdateProfile(v); err != nil {
		return err
	}
	return nil
}
