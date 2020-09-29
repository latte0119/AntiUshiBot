package main

import (
	"github.com/ChimeraCoder/anaconda"
)

type AntiUshiBot struct {
	apiClient *anaconda.TwitterApi
}

func NewAntiUshiBot(accessToken, accessTokenSecret, consumerKey, consumerSecret string) *AntiUshiBot {
	return &AntiUshiBot{
		apiClient: anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret),
	}
}

func (aub *AntiUshiBot) PostTweet(text string) error {
	_, err := aub.apiClient.PostTweet(text, nil)
	return err
}
