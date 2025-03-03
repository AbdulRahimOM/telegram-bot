package main

import (
	"fmt"

	"github.com/AbdulRahimOM/telegram-bot/internal/bot"
)

func main() {
	bot.InitBot()

	fmt.Println("Initializing bot...")
	bot.RunBot()
}
