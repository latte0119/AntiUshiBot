package main

import (
	"encoding/json"

	"github.com/koron/go-dproxy"
)

func FollowBack(query interface{}) {
	bodystr, err := dproxy.New(query).M("body").String()
	if err != nil {
		return
	}

	var body interface{}
	json.Unmarshal([]byte(bodystr), &body)
	val, _ := dproxy.New(body).M("follow_events").Value()
	screenName, err := dproxy.New(val).A(0).M("source").M("screen_name").String()

	if err != nil {
		return
	}

	api := GetTwitterAPI()
	api.FollowUser(screenName)
}
