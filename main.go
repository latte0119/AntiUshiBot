package main

import (
	"github.com/aws/aws-lambda-go/lambda"
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
	Init()

	lambda.Start(aub)
}
