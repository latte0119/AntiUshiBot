package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/latte0119/owmjp"
)

func WeatherTweet() {
	wapi := owmjp.NewAPIWithKey(os.Getenv("OWM_KEY"))
	v := url.Values{}
	v.Set("units", "metric")
	wi, err := wapi.GetCurrentWeatherData("Tokyo", v)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", wi.Weather[0].Main)

	var text string
	if tmp, err := wi.MainJP(); err == nil {
		text += fmt.Sprintf("天気:%v\n", tmp)
	}
	text += fmt.Sprintf("詳細:%v\n", wi.Weather[0].Description)
	text += fmt.Sprintf("気温:%.1f\n", wi.Main.Temp)
	text += fmt.Sprintf("湿度:%v\n", wi.Main.Humidity)
	fmt.Println(text)
	api := GetTwitterAPI()
	api.PostTweet("天気たぷにきあくん(笑顔)\n"+text, nil)
}

func ForecastTweet() {
	wapi := owmjp.NewAPIWithKey(os.Getenv("OWM_KEY"))
	v := url.Values{}
	v.Set("units", "metric")
	wis, err := wapi.GetForecastData("Tokyo", v)
	if err != nil {
		log.Fatal(err)
	}

	var text string
	for i := 0; i < 5; i++ {
		text += fmt.Sprintf("%s : %s\n", wis[7+i*8].DtTxt, wis[7+i*8].Weather[0].Description)
	}
	log.Print(text)

	api := GetTwitterAPI()
	api.PostTweet(text, nil)
}
