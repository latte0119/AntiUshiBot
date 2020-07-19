package main

import (
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

/*
Init is the initialization process
*/
func Init() {
	godotenv.Load("envfiles/.env")
	t := time.Now()
	rand.Seed(t.Unix())
}
