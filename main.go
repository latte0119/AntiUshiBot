package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"

	"github.com/aws/aws-lambda-go/lambda"
)

type User struct {
	ID         int    `json:"id"`
	ScreenName string `json:"screen_name"`
}

type Query struct {
	Mode string `json:"mode"`
	Body struct {
		FollowEvents []struct {
			Type   string `json:"type"`
			Target User   `json:"target"`
			Source User   `json:"source"`
		} `json:"follow_events"`
	} `json:"body"`
}

func aub(query Query) {
	for i := 0; i < len(query.Body.FollowEvents); i++ {
		follower := query.Body.FollowEvents[i].Source.ScreenName
		api := GetTwitterAPI()
		t := time.Now()
		api.PostTweet(fmt.Sprintf("%v %s", t.Unix(), follower), nil)
	}

	switch query.Mode {
	case "regular_tweet":
		RegularTweet()
	case "set_icon_by_time":
		SetIconByTime()
	case "weather_tweet":
		WeatherTweet()
	default:
		log.Println("undefined mode")
	}

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
