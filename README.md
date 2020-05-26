# Stock Lookup Discord Bot

Power a Discord bot which looks up and displays current data for stock symbols using Alpha Vantage. (https://www.alphavantage.co)

# Installation
    go install github.com/timkelleher/discord-bot-stocks

## Requirements

 - Golang 1.14
 - Discord Token (https://discord.com/developers/applications)
 - Alpha Vantage API Key (https://www.alphavantage.co/support/#api-key)


## Set environmental variables

Define environmental variables prior to runtime:

    export DISCORD_TOKEN=?
    export APHAVANTAGE_API_KEY=?

Or at runtime:

    DISCORD_TOKEN=? ALPHAVANTAGE_API_KEY=? go run github.com/timkelleher/discord-bot-stocks/cmd/bot

To build and execute a binary:

    go build -o stocks github.com/timkelleher/discord-bot-stocks/cmd/bot
    ./stocks

# Usage
Invite your bot to a server or direct message the bot.  The bot will scan for messages in the following format and ignore all other messages:

    stock SYMBOL

The response will be as follows:

Symbol **AAPL**
*Last Updated: 2020-05-22*
Open: 315.7700 | High: 319.2300 | Low: 315.3500 | Close: 318.8900 | Volume: 20450754
