package main

import (
	"github.com/koron/go-dproxy"
)

func FollowBack(body interface{}) error {
	val, _ := dproxy.New(body).M("follow_events").Value()
	screenName, err := dproxy.New(val).A(0).M("source").M("screen_name").String()

	if err != nil {
		return err
	}

	api := GetTwitterAPI()
	if _, err := api.FollowUser(screenName); err != nil {
		return err
	}
	return nil
}
