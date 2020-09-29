package main

import "testing"

func TestGetDDB(t *testing.T) {
	if _, err := getDDB(); err != nil {
		t.Fatal(err)
	}
}

func TestFetchTweetList(t *testing.T) {
	if _, err := fetchTweetList(); err != nil {
		t.Fatal(err)
	}
}
