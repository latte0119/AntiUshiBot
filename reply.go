package main

import (
	"fmt"

	"github.com/koron/go-dproxy"
)

func Reply(body interface{}) error {
	val, err := dproxy.New(body).M("tweet_create_events").Value()
	if err != nil {
		return err
	}

	for _, tw := range val.([]interface{}) {
		if err := PostDMToMe(fmt.Sprintf("%#v", tw)); err != nil {
			PostDMToMe(err.Error())
		}
		if _, err = dproxy.New(tw).M("in_reply_to_status_id_str").String(); err != nil {
			continue
		}
		text, _ := dproxy.New(tw).M("text").String()
		PostDMToMe(text)
	}
	return nil
}
