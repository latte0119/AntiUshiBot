package main

import (
	"testing"
)

func TestRegularTweet(t *testing.T) {
	Init()
	if err := RegularTweet(); err != nil {
		t.Fatal(err)
	}
}
