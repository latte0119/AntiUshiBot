package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/*
WeatherInfo weather info
*/
type WeatherInfo struct {
	Coord struct {
		Lat float32 `json:"lat"`
		Lon float32 `json:"lon"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"Icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Rain struct {
		Rain1h float64 `json:"rain.1h"`
		Rain3h float64 `json:"rain.3h"`
	} `json:"rain"`
	Snow struct {
		Snow1h float64 `json:"snow.1h"`
		Snow3h float64 `json:"snow.3h"`
	} `json:"snow"`
	Dt   int    `json:"dt"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
GetWeatherInfo gets weather info
*/
func GetWeatherInfo() WeatherInfo {
	baseURL := "http://api.openweathermap.org/data/2.5/weather"
	v := url.Values{}
	v.Set("q", "Tokyo")
	v.Set("units", "metric")
	v.Set("appid", "d1dcf4b83f0b4d9a5e0dbc0a059028a2")

	res, err := http.Get(baseURL + "?" + v.Encode())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherInfo WeatherInfo
	if err := json.Unmarshal(body, &weatherInfo); err != nil {
		log.Fatal(err)
	}

	return weatherInfo
}

/*
Dev uku
*/
func Dev() {
	wi := GetWeatherInfo()

	fmt.Printf("%v\n", wi.Weather[0].Main)

	var text string
	text += fmt.Sprintf("天気:%v\n", wi.Weather[0].Description)
	text += fmt.Sprintf("気温:%v\n", wi.Main.Temp)
	text += fmt.Sprintf("湿度:%v\n", wi.Main.Humidity)
	text += fmt.Sprintf("雨量%v\n", wi.Rain.Rain1h)
	fmt.Println(text)
	//api := GetTwitterAPI()
	//api.PostTweet("天気たぷにきあくん笑\n"+text, nil)
}
