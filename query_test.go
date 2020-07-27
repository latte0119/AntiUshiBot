package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func TestPostDMToMe(ti *testing.T) {

	godotenv.Load("envfiles/.env")
	PostDMToMe("latte")
}

func TestQuery(ti *testing.T) {
	godotenv.Load("envfiles/.env")

	t := time.Now()
	rand.Seed(t.Unix())
	in := []byte(`{"latte":"malta","uku":"chang"	}`)
	var query interface{}
	json.Unmarshal(in, &query)

	//str := fmt.Sprintf("%#v", query)

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	values := url.Values{}
	values.Add("description", "ukuchang")

	req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/account/update_profile.json"+"?"+values.Encode(), nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Status)
	ti.Fatal(res)
}
