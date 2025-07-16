package alphavantage

import (
	"encoding/json"
	"fmt"
	"time"
)

// OptionContract represents a single option contract (either call or put) with detailed option Greeks and other metrics.
type OptionContract struct {
	ContractID        string    `json:"contractID"`
	Symbol            string    `json:"symbol"`
	Expiration        string    `json:"expiration"`
	Strike            AVFloat64 `json:"strike"`
	Type              string    `json:"type"` // "call" or "put"
	Last              AVFloat64 `json:"last"`
	Mark              AVFloat64 `json:"mark"`
	Bid               AVFloat64 `json:"bid"`
	BidSize           AVInt     `json:"bid_size"`
	Ask               AVFloat64 `json:"ask"`
	AskSize           AVInt     `json:"ask_size"`
	Volume            AVInt     `json:"volume"`
	OpenInterest      AVInt     `json:"open_interest"`
	Date              string    `json:"date"`
	ImpliedVolatility AVFloat64 `json:"implied_volatility"`
	Delta             AVFloat64 `json:"delta"`
	Gamma             AVFloat64 `json:"gamma"`
	Theta             AVFloat64 `json:"theta"`
	Vega              AVFloat64 `json:"vega"`
	Rho               AVFloat64 `json:"rho"`
}

// HistoricalOptionsData represents the full historical options data for a given symbol.
type HistoricalOptionsData struct {
	Symbol   string           `json:"symbol"`
	Endpoint string           `json:"endpoint"`
	Message  string           `json:"message"`
	Data     []OptionContract `json:"data"`
}

// toHistoricalOptionsData parses the JSON response into the HistoricalOptionsData struct.
func toHistoricalOptionsData(buf []byte) (*HistoricalOptionsData, error) {
	optionsData := &HistoricalOptionsData{}
	if err := json.Unmarshal(buf, optionsData); err != nil {
		return nil, err
	}
	return optionsData, nil
}

// HistoricalOptions fetches and returns the historical options data for the specified company symbol.
func (c *Client) HistoricalOptions(symbol string, date *time.Time) (*HistoricalOptionsData, error) {
	const function = "HISTORICAL_OPTIONS"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)

	// Add date parameter if provided
	if date != nil {
		url = fmt.Sprintf("%s&date=%s", url, date.Format("2006-01-02"))
	}

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toHistoricalOptionsData(body)
}
