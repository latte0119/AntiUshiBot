package main

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestForecast(t *testing.T) {
	godotenv.Load("envfiles/.env")

	ForecastTweet()
}
