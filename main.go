package main

import (
	"flag"
	"fmt"
	"wetherForecast/geo"
	"wetherForecast/wether"
)

func main() {
	fmt.Println("New project")
	city := flag.String("city", "", "user`s city")
	format := flag.Int("format", 1, "wether output format")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(geoData)

	wetherData := wether.GetWether(*geoData, *format)
	fmt.Println(wetherData)
}
