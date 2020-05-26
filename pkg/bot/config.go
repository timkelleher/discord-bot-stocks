package bot

import (
	"os"

	"github.com/timkelleher/discord-bot-stocks/pkg/stock"
)

type Config struct {
	BotConfig   BotConfig
	StockConfig stock.StockConfig
}

type BotConfig struct {
	DiscordToken string
}

func (c *Config) FillFromEnv() {
	c.BotConfig.DiscordToken = os.Getenv("DISCORD_TOKEN")
	c.StockConfig.AlphaVantageAPIKey = os.Getenv("ALPHAVANTAGE_API_KEY")
}
