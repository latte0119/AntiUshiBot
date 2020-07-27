package main

import (
	"testing"
)

func TestWeatherTweet(t *testing.T) {
	Init()
	WeatherTweet()
}

func TestUpdateNameWithEmoji(t *testing.T) {
	Init()
	UpdateNameWithEmoji("アンチうしbot")
}
