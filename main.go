package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load("envfiles/.env")

	t := time.Now()
	rand.Seed(t.Unix())

	mode := flag.String("mode", "regular_tweet", "mode")
	flag.Parse()

	switch *mode {
	case "regular_tweet":
		RegularTweet()
	default:
		fmt.Println("undefined mode")
	}

}
