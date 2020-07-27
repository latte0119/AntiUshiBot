package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestTimeTweet(te *testing.T) {
	godotenv.Load("envfiles/.env")

	t := time.Now()
	rand.Seed(t.Unix())

	TimeTweet()
}
