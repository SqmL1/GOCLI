package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name    string `json:name`
		Country string `json:country`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:temp_c`
		Condition struct {
			Text string `json:text`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				Timeepoch int64   `json:time_epoch`
				TempC     float64 `json:temp_c`
				Condition struct {
					Text string `json:text`
				} `json:""`
				Chancerain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=40d8184d067b46af963151231241205&q=Blaine&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	fmt.Println(weather)

	location, current, _ := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s %s %0.f %s",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
}
