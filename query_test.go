package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestQuery(ti *testing.T) {
	godotenv.Load("envfiles/.env")

	t := time.Now()
	rand.Seed(t.Unix())
	in := []byte(`{"latte":"malta","uku":"chang"	}`)
	ti.Fatal(string(in))
	var query interface{}
	json.Unmarshal(in, &query)

	str := fmt.Sprintf("%#v", query)

	api := GetTwitterAPI()
	_, err := api.PostTweet(str, nil)
	if err != nil {
		ti.Fatal("uuuu")
	}
}
