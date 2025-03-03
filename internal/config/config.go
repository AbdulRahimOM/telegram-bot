package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	WeatherApiKey    string
	TelegramBotToken string
)

func init() {
	LoadEnv()
}

func LoadEnv() {
	fmt.Println("Loading .env file...")
	//parse .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. err", err)
	}

	TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	if TelegramBotToken == "" {
		log.Fatal("Bot token is missing!")
	}

	WeatherApiKey = os.Getenv("OPENWEATHER_API_KEY")
	if WeatherApiKey == "" {
		log.Fatal("Weather API key is missing!")
	}

}
