package main

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestWeatherTweet(t *testing.T) {
	godotenv.Load("envfiles/.env")
	WeatherTweet()
}
