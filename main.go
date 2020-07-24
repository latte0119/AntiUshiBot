package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/aws/aws-lambda-go/lambda"
)

type Query struct {
	Mode string `json:"mode"`
}

func aub(query Query) {
	switch query.Mode {
	case "regular_tweet":
		RegularTweet()
	case "set_icon_by_time":
		SetIconByTime()
	case "many_tweets":
		ManyTweets()
	case "dev":
		Dev()
	default:
		log.Println("undefined mode")
	}

}

func main() {
	godotenv.Load("envfiles/.env")

	t := time.Now()
	rand.Seed(t.Unix())

	//mode := flag.String("mode", "", "mode")
	//flag.Parse()
	lambda.Start(aub)
}
