package bot

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/timkelleher/discord-bot-stocks/pkg/stock"
)

var config Config
var stockRegex = regexp.MustCompile(`(?i)^stock (\w+)$`)

func Run(c Config) {
	config = c

	discord, err := discordgo.New("Bot " + config.BotConfig.DiscordToken)
	if err != nil {
		fmt.Println("Error creating Discord session", err)
		return
	}

	// Register the MessageCreate func as a callback for MessageCreate events.
	discord.AddHandler(MessageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("[%s] %s: %s\n", m.ChannelID, m.Author.Username, m.Content)

	symbol := ExtractStockSymbol(m.Content)
	if symbol != "" {
		stockData := stock.Fetch(config.StockConfig, symbol)
		s.ChannelMessageSend(m.ChannelID, stockData.String())
		fmt.Println("Response: ", stockData.String())
	}
}

func ExtractStockSymbol(input string) string {
	var symbol string
	if stockRegex.MatchString(input) {
		request := input
		words := strings.Fields(request)
		if len(words) > 1 && words[1] != "" {
			symbol = words[1]
		}
	}

	return symbol
}
