package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dghubble/oauth1"
)

func SetIconByTime() error {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(loc)

	k := t.Hour() % 12

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	file, _ := os.Open(fmt.Sprintf("image/rotate/icon%d.jpg", k))
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	en := base64.StdEncoding.EncodeToString(data)

	values := url.Values{}
	values.Add("image", en)

	req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/account/update_profile_image.json"+"?"+values.Encode(), nil)
	if err != nil {
		return err
	}
	if _, err = httpClient.Do(req); err != nil {
		return err
	}
	return nil
}
