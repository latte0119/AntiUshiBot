package main

import (
	"fmt"
	"net/url"
	//"time"
)

func main() {
	api := GetTwitterAPI()

	/*
		t := time.Now()
		tweet, err := api.PostTweet("うしたぷにきあくん楽しみ〜"+t.String(), nil)
	*/

	v := url.Values{}
	v.Set("screen_name", "_ei13333")
	timeline, _ := api.GetUserTimeline(v)

	for _, tw := range timeline {
		fmt.Println(tw.FullText)
	}

	dest := timeline[0].IdStr

	v = url.Values{}
	v.Set("in_reply_to_status_id", dest)
	v.Set("auto_populate_reply_metadata", "true")
	api.PostTweet("でも君，う し た ぷ に き あ く ん 笑 じゃん", v)
}
