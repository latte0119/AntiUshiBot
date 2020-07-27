package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestRegularTweet(t *testing.T) {
	godotenv.Load(".env")
	rand.Seed(time.Now().Unix())

	RegularTweet()
}
