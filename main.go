package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koron/go-dproxy"
)

func aub(query interface{}) {
	bodystr, _ := dproxy.New(query).M("body").String()
	var body interface{}
	json.Unmarshal([]byte(bodystr), &body)

	FollowBack(body)
	Reply(body)

	mode, _ := dproxy.New(query).M("mode").String()

	switch mode {
	case "regular_tweet":
		RegularTweet()
	case "set_icon_by_time":
		SetIconByTime()
	case "update_name_with_emoji":
		UpdateNameWithEmoji("アンチうしbot")
	default:
	}
}

func main() {
	Init()

	lambda.Start(aub)
}
