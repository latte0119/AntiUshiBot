package main

import (
	"fmt"
	"testing"
)

func TestGenWeatherTweet(t *testing.T) {
	Init()
	text, err := GenWeatherTweet()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(text)
}

func TestGenForecastTweet(t *testing.T) {
	Init()
	text, err := GenForecastTweet()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(text)
}

func TestGenForecastEmojiTweet(t *testing.T) {
	Init()
	text, err := GenForecastEmojiTweet()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(text)
}

func TestUpdateNameWithEmoji(t *testing.T) {
	Init()
	if err := UpdateNameWithEmoji("アンチうしbot"); err != nil {
		t.Fatal(err)
	}
}
