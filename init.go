package main

import (
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(".env")
	rand.Seed(time.Now().Unix())
}
