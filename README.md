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

## Detailed Setup Guide

### 1. Get a Telegram Bot Token
- Open Telegram and search for [@BotFather](https://t.me/botfather)
- Start a chat and use `/newbot` command
- Follow the prompts to create your bot
- Copy the provided bot token
- Note: Keep your token secure and never share it publicly

### 2. Get a Weather API Key
- Sign up at [OpenWeatherMap](https://openweathermap.org/)
- Subscribe to the free plan
- Navigate to API keys section
- Copy your API key
- Note: API key activation may take a few hours

### 3. Set Up Environment Variables
Create a `.env` file in the project root:
```bash
# .env
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
WEATHER_API_KEY=your_openweather_api_key
```

### 4. Run the Bot
```bash
# Start the bot
go run cmd/main.go

# You should see output indicating successful connection
```

### 5. Use the Bot
- Open Telegram and search for your bot username
- Start a conversation with `/start`
- Try sharing your location or using commands like:
  ```
  /help - View all commands
  /setlocation(12.345,67.890) - Set default location
  /info - Get weather for default location
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

## Code Structure

### Main Components

- **cmd/main.go**: Entry point of the application that initializes and runs the bot
- **internal/bot/bot.go**: Core bot logic including:
  - Message handling
  - Command processing
  - Rate limiting implementation
  - User session management
- **internal/weather_app/weather.go**: Weather service integration
  - OpenWeather API client
  - Weather data structures
  - API response handling
- **internal/config/**: Configuration management for API keys

### Key Features Implementation

- **Rate Limiting**: Uses `golang.org/x/time/rate` with per-user limiters stored in a sync.Map
- **Location Management**: Maintains user default locations in memory using a map
- **Weather Data**: Structured using custom types for clean API response handling
- **Error Handling**: Comprehensive error handling for API calls and user inputs

## Weather Information Provided
- Temperature (°C)
- Humidity (%)
- Atmospheric Pressure (hPa)
- Wind Speed (m/s)
- Timezone information
