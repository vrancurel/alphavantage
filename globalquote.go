package alphavantage

import (
	"encoding/json"
	"fmt"
	// "github.com/go-playground/validator/v10"
)

// GlobalQuoteResponse - encapsulates global quote repsonse
type GlobalQuoteResponse struct {
	GlobalQuote GlobalQuote `json:"Global Quote"`
}

// GlobalQuote - encapsulates global quote
type GlobalQuote struct {
	Symbol           string    `json:"01. symbol" validate:"required"`
	Open             float64   `json:"02. open,string"`
	High             float64   `json:"03. high,string"`
	Low              float64   `json:"04. low,string"`
	Price            float64   `json:"05. price,string"`
	Volume           int       `json:"06. volume,string"`
	LatestTradingDay string    `json:"07. latest trading day"`
	PreviousClose    float64   `json:"08. previous close,string"`
	Change           float64   `json:"09. change,string"`
	ChangePercent    AVPercent `json:"10. change percent"`
}

func toGlobalQuote(buf []byte) (*GlobalQuote, error) {
	globalQuoteResponse := &GlobalQuoteResponse{}
	if err := json.Unmarshal(buf, globalQuoteResponse); err != nil {
		return nil, err
	}

	// validation is a nice feature but it can break if some
	// symbols have no data (we could have a strict mode)
	// validate := validator.New()
	// err := validate.Struct(globalQuoteResponse.GlobalQuote)
	// if err != nil {
	// return nil, err
	// }

	return &globalQuoteResponse.GlobalQuote, nil
}

// GlobalQuote fetches data from the Global Quote endpoint for the given symbol
func (c *Client) GlobalQuote(symbol string) (*GlobalQuote, error) {
	const functionName = "GLOBAL_QUOTE"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s",
		baseURL, functionName, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	globalQuote, err := toGlobalQuote(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return globalQuote, nil
}
