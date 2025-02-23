package weather

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetWeather(location string) (WeatherData, error) {
	resp, err := http.Get("https://wttr.in/" + location + "?format=j1")
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, err
	}

	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return WeatherData{}, err
	}

	return weather, nil
}
