package main

import (
	"os"
	"testing"
)

func TestPostTweet(t *testing.T) {
	t.Skip()
	Init()
	antiUshiBot := NewAntiUshiBot(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"))
	if err := antiUshiBot.PostTweet("test"); err != nil {
		t.Fatal(err)
	}
}
