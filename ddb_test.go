package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestDDB(t *testing.T) {
	godotenv.Load("envfiles/.env")

	ti := time.Now()
	rand.Seed(ti.Unix())
	RegularTweet()
}
