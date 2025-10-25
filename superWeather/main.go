package main

import (
	"flag"
	"fmt"
	"super-weather/geo"
)

func main() {
	city := flag.String("city", "", "город пользователя")
	flag.Parse()

	fmt.Println(*city)



	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

		fmt.Println(geoData)
}