# Telegram Weather Bot

A Telegram bot that provides real-time weather information. Simply share your location or set default coordinates to get instant weather updates including temperature, humidity, pressure, and wind speed.

## Features

- 🌍 Get weather by sharing location
- 📍 Set default location using coordinates
- 🌡️ Real-time weather metrics
- ⚡ Fast response time
- 🛡️ Rate limiting protection

## Bot Commands

| Command | Description |
|---------|------------|
| `/start` | Welcome message |
| `/help` | Show available commands and usage |
| `/info` | Get weather for default location |
| `/setlocation(lat,lon)` | Set default location |
| `/status` | Check bot status |

Example: `/setlocation(12.345,67.890)`

## Setup

### Prerequisites
- Go programming language
- Telegram Bot Token (from [@BotFather](https://t.me/botfather))
- OpenWeather API Key (from [OpenWeather](https://openweathermap.org/api))

### Environment Variables
```bash
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
WEATHER_API_KEY=your_openweather_api_key
```

### Quick Start
1. Clone the repository
```bash
git clone https://github.com/AbdulRahimOM/telegram-bot.git
cd telegram-bot
```

2. Set environment variables
3. Run the bot
```bash
go run cmd/main.go
```

## Project Structure
```
.
├── cmd/
│   └── main.go          # Entry point
├── internal/
│   ├── bot/            # Bot logic
│   ├── config/         # Configuration
│   └── weather_app/    # Weather API integration
└── README.md
```

## Rate Limiting
- 10 requests per minute per user
- Helps prevent API abuse
- User-specific rate limiting

## Weather Information Provided
- Temperature (°C)
- Humidity (%)
- Atmospheric Pressure (hPa)
- Wind Speed (m/s)
- Timezone information
