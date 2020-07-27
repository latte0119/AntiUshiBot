package main

import (
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/koron/go-dproxy"
)

func aub(query interface{}) {
	FollowBack(query)

	mode, _ := dproxy.New(query).M("mode").String()

	switch mode {
	case "regular_tweet":
		RegularTweet()
	case "set_icon_by_time":
		SetIconByTime()
	default:
	}
}

func main() {
	godotenv.Load(".env")
	rand.Seed(time.Now().Unix())

	lambda.Start(aub)
}
