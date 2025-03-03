package weatherapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbdulRahimOM/telegram-bot/internal/config"
)

const apiURL = "https://api.openweathermap.org/data/3.0/onecall"

type WeatherResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Temp      float64 `json:"temp"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
		WindSpeed float64 `json:"wind_speed"`
		Weather   []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"current"`
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func GetWeather(lat, lon float64) (*WeatherResponse, error) {
	apiKey := config.WeatherApiKey

	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&exclude=minutely,hourly,daily&units=metric", apiURL, lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var weatherData WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
