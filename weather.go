package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/latte0119/owmjp"
)

func GenWeatherTweet() (string, error) {
	wapi := owmjp.NewAPIWithKey(os.Getenv("OWM_KEY"))
	v := url.Values{}
	v.Set("units", "metric")
	wi, err := wapi.GetCurrentWeatherData("Tokyo", v)

	if err != nil {
		return "", err
	}

	var text string
	if tmp, err := wi.MainJP(); err == nil {
		text += fmt.Sprintf("天気:%v\n", tmp)
	}
	text += fmt.Sprintf("詳細:%v\n", wi.Weather[0].Description)
	text += fmt.Sprintf("気温:%.1f\n", wi.Main.Temp)
	text += fmt.Sprintf("湿度:%v\n", wi.Main.Humidity)
	return text, nil
}

func GenForecastTweet() (string, error) {
	wapi := owmjp.NewAPIWithKey(os.Getenv("OWM_KEY"))
	v := url.Values{}
	v.Set("units", "metric")
	wis, err := wapi.GetForecastData("Tokyo", v)
	if err != nil {
		return "", err
	}

	var text string
	for i := 0; i < 5; i++ {
		text += fmt.Sprintf("%s : %s\n", wis[7+i*8].DtTxt, wis[7+i*8].Weather[0].Description)
	}

	return text, nil
}

func GenForecastEmojiTweet() (string, error) {
	wapi := owmjp.NewAPIWithKey(os.Getenv("OWM_KEY"))
	wis, err := wapi.GetForecastData("Tokyo", nil)
	if err != nil {
		return "", err
	}

	var text string
	for i, w := range wis {
		e, err := w.Emoji()
		if err != nil {
			return "", err
		}
		text += e
		if i%8 == 7 {
			text += "\n"
		}
	}
	return text, nil
}

func UpdateNameWithEmoji(orig string) error {
	wapi := owmjp.NewAPIWithKey(os.Getenv("OWM_KEY"))
	w, err := wapi.GetCurrentWeatherData("Tokyo", nil)
	if err != nil {
		return err
	}

	e, err := w.Emoji()
	if err != nil {
		return err
	}
	if err := UpdateName(e + orig + e); err != nil {
		return err
	}
	return nil
}
