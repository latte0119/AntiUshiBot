package main

import (
	"fmt"
	"time"
)

func main() {
	api := GetTwitterAPI()

	t := time.Now()
	tweet, err := api.PostTweet("うしたぷにきあくん楽しみ〜"+t.String(), nil)

	if err == nil {
		fmt.Println(tweet.Text)
	} else {
		fmt.Println("uku")
	}
}
