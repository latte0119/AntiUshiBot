package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/image/math/f64"

	"github.com/ChimeraCoder/anaconda"
	"golang.org/x/image/draw"
)

func EscapeProcessing(text string) string {
	text = strings.Replace(text, "\\\"", "\"", -1)
	//text = strings.Replace(text, "\\", "\\\\", -1)
	//	text = strings.Replace(text, "\\", "X", -1)
	text = strings.Replace(text, "\"", "\\\"", -1)
	text = strings.Replace(text, "/", "\\/", -1)
	return text
}

func GenerateRotatedIcons() {

	fileIn, err := os.Open("image/icon.png")

	if err != nil {
		log.Fatal(err)
	}

	defer fileIn.Close()

	img, err := png.Decode(fileIn)

	if err != nil {
		log.Fatal(err)
	}

	for k := 0; k < 12; k++ {
		out := image.NewRGBA(img.Bounds())

		deg := 30.0 * float64(k)
		rad := deg * math.Pi / 180

		cx := float64(img.Bounds().Min.X+img.Bounds().Max.X) / 2.0
		cy := float64(img.Bounds().Min.Y+img.Bounds().Max.Y) / 2.0

		cos, sin := math.Cos(rad), math.Sin(rad)

		ax := -(cos*cx - sin*cy) + cx
		ay := -(sin*cx + cos*cy) + cy
		t := f64.Aff3{
			cos, -sin, ax,
			sin, cos, ay,
		}

		draw.NearestNeighbor.Transform(out, t, img, img.Bounds(), draw.Over, nil)

		fileOut, err := os.Create(fmt.Sprintf("image/rotate/icon%d.jpg", k))

		if err != nil {
			log.Fatal(err)
		}
		defer fileOut.Close()

		if err := jpeg.Encode(fileOut, out, nil); err != nil {
			log.Fatal(err)
		}
	}

}

func GetDetailedTime() string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(loc)

	return t.Format(time.ANSIC)
}

func deleteAll(api *anaconda.TwitterApi) {
	v := url.Values{}
	v.Set("screen_name", "_ei133333")
	timeline, _ := api.GetUserTimeline(v)
	for _, tw := range timeline {
		api.DeleteTweet(tw.Id, false)
	}
}
