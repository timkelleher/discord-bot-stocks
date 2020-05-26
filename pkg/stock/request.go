package stock

import "fmt"

type Request struct {
	Symbol   string
	APIKey   string
	Function string
}

func (r Request) BuildUrl() string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=%s&symbol=%s&interval=5min&apikey=%s", r.Function, r.Symbol, r.APIKey)
}
