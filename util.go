package main

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func deleteAll(api *anaconda.TwitterApi) {
	v := url.Values{}
	v.Set("screen_name", "_ei133333")
	timeline, _ := api.GetUserTimeline(v)
	for _, tw := range timeline {
		api.DeleteTweet(tw.Id, false)
	}
}
