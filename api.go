package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
)

const API = "https://api.openweathermap.org/data/2.5/weather"

var WEATHER_API_KEY = os.Getenv("WEATHER_API_KEY")

func get_weather_by_coords(longitude float64, latitude float64) (string, error) {

	request := &fasthttp.Request{}
	response := &fasthttp.Response{}

	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&lang=ru&units=metric", API, latitude, longitude, WEATHER_API_KEY)
	request.Header.SetMethod(http.MethodGet)
	request.SetRequestURI(url)

	err := fasthttp.Do(request, response)
	if err != nil {
		return "😢 Не удалось получить погоду для вас, попробуйте позднее", err
	}
	var weather weatherResponse
	err = json.Unmarshal(response.Body(), &weather)

	if err != nil {
		fmt.Print(err)
		return "", err
	}
	emodji := getEmodji(weather.Weather[0].Icon)
	result := fmt.Sprintf(
		"☂️ Погода в городе %s:\n%s %s\n🌡 Температура: %.f °C\n🌡 По ощущениям: %.f °C\n",
		weather.City, emodji, weather.Weather[0].Description, weather.Main.Temp, weather.Main.FeelLikes)

	return result, nil
}

func getEmodji(icon string) string {
	var emodjiMap = map[string]string{
		"01": "☀️",
		"02": "⛅️",
		"03": "☁️",
		"04": "☁️",
		"09": "🌧",
		"10": "🌦",
		"11": "⛈",
		"13": "❄️",
		"50": "🌫",
	}
	return emodjiMap[icon[:len(icon)-1]]
}

type weatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelLikes float64 `json:"feels_like"`
	} `json:"main"`
	City string `json:"name"`
}
