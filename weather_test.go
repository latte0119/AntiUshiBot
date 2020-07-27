package main

import (
	"testing"
)

func TestWeatherTweet(t *testing.T) {
	Init()
	if err := WeatherTweet(); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateNameWithEmoji(t *testing.T) {
	Init()
	if err := UpdateNameWithEmoji("アンチうしbot"); err != nil {
		t.Fatal(err)
	}
}
