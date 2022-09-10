package main

import (
	"fmt"
	"time"

	"github.com/kelvins/sunrisesunset"
)

func main() {
	params := sunrisesunset.Parameters{
		Latitude:  35.70012,
		Longitude: 139.79013,
		UtcOffset: +9.0,
		Date:      time.Date(2022, 9, 13, 0, 0, 0, 0, time.UTC),
	}

	sunrise, sunset, err := params.GetSunriseSunset()

	if err == nil {
		fmt.Println("Sunrise:", sunrise.Format("15:04:05"))
		fmt.Println("Sunset:", sunset.Format("15:04:05"))
	} else {
		fmt.Println(err)
	}
}
