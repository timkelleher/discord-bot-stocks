package main

import (
	"github.com/timkelleher/discord-bot-stocks/pkg/bot"
)

func main() {
	var config bot.Config
	config.FillFromEnv()
	bot.Run(config)
}
