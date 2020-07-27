package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/joho/godotenv"

	"github.com/aws/aws-lambda-go/lambda"
)

type User struct {
	ID         int    `json:"id"`
	ScreenName string `json:"screen_name"`
}

func aub(query interface{}) {
	str := fmt.Sprintf("%#v", query)
	api := GetTwitterAPI()
	api.PostDMToScreenName(str, "latte0119_")
}

func TimeTweet() {
	api := GetTwitterAPI()
	t := time.Now()
	api.PostTweet(t.String(), nil)
}

func main() {
	godotenv.Load("envfiles/.env")

	t := time.Now()
	rand.Seed(t.Unix())

	//mode := flag.String("mode", "", "mode")
	//flag.Parse()
	lambda.Start(aub)
}
