package stock

import (
	"fmt"
	"strings"
)

type Response struct {
	Metadata Metadata             `json:"Meta Data"`
	Data     map[string]StockData `json:"Time Series (Daily)"`
}

type Metadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type StockData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

/*
 * Metadata.LastRefreshed is in either YYYY-MM-DD HH-II-SS or YYYY-MM-DD format
 * Use last refreshed's date for key to map[string]StockData
 */
func (r Response) LastRefreshedDate() string {
	dateSlice := strings.Fields(r.Metadata.LastRefreshed)
	return dateSlice[0]
}

func (r Response) String() string {
	if r.Metadata.Symbol == "" {
		return "No data found for symbol"
	}

	symbol := strings.ToUpper(r.Metadata.Symbol)
	stockData := r.Data[r.LastRefreshedDate()]
	if stockData.Open == "" {
		return fmt.Sprintf("Error fetching data for %s", symbol)
	}

	return fmt.Sprintf("Symbol **%s**\n*Last Updated: %s*\nOpen: %s | High: %s | Low: %s | Close: %s | Volume: %s\n", symbol, r.Metadata.LastRefreshed, stockData.Open, stockData.High, stockData.Low, stockData.Close, stockData.Volume)
}
