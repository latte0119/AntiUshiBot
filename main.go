package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/koron/go-dproxy"
)

func aub(query interface{}) {
	bodystr, err := dproxy.New(query).M("body").String()
	if err != nil {
		PostDMToMe(err.Error())
		return
	}
	var body interface{}
	json.Unmarshal([]byte(bodystr), &body)
	val, err := dproxy.New(body).M("follow_events").Value()
	str, err := dproxy.New(val).A(0).M("source").M("screen_name").String()
	if err == nil {
		api := GetTwitterAPI()
		api.PostTweet(fmt.Sprintf("followed by %s", str), nil)
		//PostDMToMe(fmt.Sprintf("%#v", str))
	} else {
		//PostDMToMe(err.Error())
	}
}

func main() {
	godotenv.Load(".env")
	rand.Seed(time.Now().Unix())

	lambda.Start(aub)
}
