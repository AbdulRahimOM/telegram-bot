package bot

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/AbdulRahimOM/telegram-bot/internal/config"
	weatherapp "github.com/AbdulRahimOM/telegram-bot/internal/weather_app"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/time/rate"
)

var (
	bot          *tgbotapi.BotAPI
	updates      tgbotapi.UpdatesChannel
	userLimiters sync.Map

	defaultLocation = make(map[int64]weatherapp.Location)
)

const (
	greetingMessage = "Hello! Welcome to the bot"
	helpMessage     = `This is a bot that can help you with information. 
	Attach your location to get weather information.

	You can also set your default location by sending your location and typing /setlocation(latitude,longitude)
	eg: /setlocation(12.345,67.890)

	Commands:
	/start - Greet the bot
	/help - Get help
	/info - Get weather information of your default location
	/setlocation(latitude,longitude) - Set your default location

	You can also send your location by attaching it to get weather information. (use the paperclip icon->location)

	`
	dontKnowMessage = "I don't know what to say 🤷‍♂️"

	maxRequestsPerMinute = 10 //change this in rateLimitMsg constant also
	rateLimitMsg         = "Rate limit exceeded 10 requests per minute). Please wait."
)

func InitBot() {
	var err error
	fmt.Println("Bot token:", config.TelegramBotToken) ////////////

	bot, err = tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error initializing bot: %v", err))
	}

	bot.Debug = true // Enable debug logs

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates = bot.GetUpdatesChan(u)
}

func getLimiter(userID int64) *rate.Limiter {
	limiter, exists := userLimiters.Load(userID)
	if !exists {
		limiter = rate.NewLimiter(rate.Every(time.Minute/time.Duration(maxRequestsPerMinute)), maxRequestsPerMinute)
		userLimiters.Store(userID, limiter)
	}
	return limiter.(*rate.Limiter)
}

func RunBot() {
	var msg tgbotapi.MessageConfig
	fmt.Println("Running bot...")
	for update := range updates {
		if update.Message == nil {
			continue
		}

		userID := update.Message.Chat.ID
		log.Printf("[%d] %s", userID, update.Message.Text)

		limiter := getLimiter(userID)

		// Apply rate limiting
		if !limiter.Allow() {
			bot.Send(tgbotapi.NewMessage(userID, rateLimitMsg))
			continue
		}

		if update.Message.Location == nil {
			switch update.Message.Text {
			case "/start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, greetingMessage)
			case "/help":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, helpMessage)
			case "/info":
				defaultLocation := defaultLocation[update.Message.Chat.ID]
				weather, err := weatherapp.GetWeather(defaultLocation.Latitude, defaultLocation.Longitude)
				if err != nil {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Error getting weather data")
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(
						`Weather in your default location: 
						(Longitude:%f,Latitude:%f) 
						Timezone%s: 
						Temperature: %.2f°C, 
						%d%% humidity, 
						%d hPa pressure, 
						%.2f m/s wind speed`,
						defaultLocation.Latitude, defaultLocation.Longitude, weather.Timezone, weather.Current.Temp, weather.Current.Humidity, weather.Current.Pressure, weather.Current.WindSpeed))
				}
			case "/status":
				defaultLocation := defaultLocation[update.Message.Chat.ID]
				if defaultLocation.Latitude == 0 && defaultLocation.Longitude == 0 {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, `Bot is running.
						Default location is not set.`)
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(`Bot is running.
						Default location is set to (%f, %f)`, defaultLocation.Latitude, defaultLocation.Longitude))
				}
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Bot is running")
			default:
				if strings.HasPrefix(update.Message.Text, "/setlocation") {
					// Parse latitude and longitude
					var lat, lon float64
					_, err := fmt.Sscanf(update.Message.Text, "/setlocation(%f,%f)", &lat, &lon)
					if err != nil {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid location format")
					}
					defaultLocation[update.Message.Chat.ID] = weatherapp.Location{Latitude: lat, Longitude: lon}
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Default location set to (%f, %f)", lat, lon))
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, dontKnowMessage)
				}
			}
		} else {
			lat := update.Message.Location.Latitude
			lon := update.Message.Location.Longitude

			// Fetch weather data
			weatherInfo, err := weatherapp.GetWeather(lat, lon)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Error getting weather data")
			}
			// text := fmt.Sprintf("Weather in %s: %.2f°C, %d%% humidity, %d hPa pressure, %.2f m/s wind speed", weatherInfo.Timezone, weatherInfo.Current.Temp, weatherInfo.Current.Humidity, weatherInfo.Current.Pressure, weatherInfo.Current.WindSpeed)
			text := fmt.Sprintf(`Weather in your current location:
			(Longitude:%f,Latitude:%f)
			Timezone%s:
			Temperature: %.2f°C,
			%d%% humidity,
			%d hPa pressure,
			%.2f m/s wind speed`,
				lat, lon, weatherInfo.Timezone, weatherInfo.Current.Temp, weatherInfo.Current.Humidity, weatherInfo.Current.Pressure, weatherInfo.Current.WindSpeed)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
		bot.Send(msg)
	}
}
