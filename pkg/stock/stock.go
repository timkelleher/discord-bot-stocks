package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StockApi struct {
	Config StockConfig
}

func Fetch(config StockConfig, symbol string) Response {
	var request Request
	request.Symbol = symbol
	request.APIKey = config.AlphaVantageAPIKey
	request.Function = "TIME_SERIES_DAILY"
	response := QueryDailyApi(request)
	return response
}

func QueryDailyApi(r Request) Response {

	fmt.Println("URL: %s", r.BuildUrl())

	resp, err := http.Get(r.BuildUrl())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return response
}
