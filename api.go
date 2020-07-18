package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

/*
GetTwitterAPI gets the twitter API
*/
func GetTwitterAPI() *anaconda.TwitterApi {

	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	return api
}
