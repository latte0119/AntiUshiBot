package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

/*
GetTwitterAPI gets the twitter API
*/
func GetTwitterAPI() *anaconda.TwitterApi {
	godotenv.Load("envfiles/.env")

	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	return api
}
