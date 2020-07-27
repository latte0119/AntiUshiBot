package main

import (
	"net/url"
	"strings"

	"github.com/koron/go-dproxy"
)

func Reply(body interface{}) error {
	forUserID, err := dproxy.New(body).M("for_user_id").String()
	if err != nil {
		return err
	}

	val, err := dproxy.New(body).M("tweet_create_events").Value()
	if err != nil {
		return err
	}

	for _, tw := range val.([]interface{}) {
		if _, err := dproxy.New(tw).M("in_reply_to_status_id_str").String(); err != nil {
			continue
		}

		if id, _ := dproxy.New(tw).M("user").M("id_str").String(); id == forUserID {
			continue
		}

		screenName, _ := dproxy.New(tw).M("user").M("screen_name").String()

		tweetID, _ := dproxy.New(tw).M("id_str").String()

		if origText, _ := dproxy.New(tw).M("text").String(); strings.Contains(origText, "天気") {
			text, err := GenWeatherTweet()
			if err != nil {
				return err
			}

			api := GetTwitterAPI()
			v := url.Values{}
			v.Add("in_reply_to_status_id", tweetID)

			api.PostTweet("@"+screenName+" "+GetDetailedTime()+"\n"+text, v)
			continue
		}

		if origText, _ := dproxy.New(tw).M("text").String(); strings.Contains(origText, "予報") {
			text, err := GenForecastTweet()
			if err != nil {
				return err
			}

			api := GetTwitterAPI()
			v := url.Values{}
			v.Add("in_reply_to_status_id", tweetID)

			api.PostTweet("@"+screenName+" "+GetDetailedTime()+"\n"+text, v)
			continue
		}

		if origText, _ := dproxy.New(tw).M("text").String(); strings.Contains(origText, "謎") {
			text, err := GenForecastEmojiTweet()
			if err != nil {
				return err
			}

			api := GetTwitterAPI()
			v := url.Values{}
			v.Add("in_reply_to_status_id", tweetID)

			api.PostTweet("@"+screenName+" "+GetDetailedTime()+"\n"+text, v)
			continue
		}

	}
	return nil
}
