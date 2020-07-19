package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load("envfiles/.env")

	t := time.Now()
	rand.Seed(t.Unix())

	mode := flag.String("mode", "", "mode")
	flag.Parse()

	switch *mode {
	case "regular_tweet":
		RegularTweet()
	case "set_icon_by_time":
		SetIconByTime()
	default:
		log.Println("undefined mode")
	}

}
