package alphavantage

import (
	"encoding/json"
	"fmt"
)

// InsiderTransaction represents a single insider transaction record.
type InsiderTransaction struct {
	TransactionDate       string `json:"transaction_date"`
	Ticker                string `json:"ticker"`
	Executive             string `json:"executive"`
	ExecutiveTitle        string `json:"executive_title"`
	SecurityType          string `json:"security_type"`
	AcquisitionOrDisposal string `json:"acquisition_or_disposal"`
	Shares                string `json:"shares"`
	SharePrice            string `json:"share_price"`
}

// InsiderTransactions represents the response from the INSIDER_TRANSACTIONS API endpoint.
type InsiderTransactions struct {
	Data []InsiderTransaction `json:"data"`
}

func toInsiderTransactions(buf []byte) (*InsiderTransactions, error) {
	insiderTransactions := &InsiderTransactions{}
	if err := json.Unmarshal(buf, insiderTransactions); err != nil {
		return nil, err
	}
	return insiderTransactions, nil
}

// InsiderTransactions fetches and returns the insider transactions data for the specified symbol.
func (c *Client) InsiderTransactions(symbol string) (*InsiderTransactions, error) {
	const function = "INSIDER_TRANSACTIONS"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toInsiderTransactions(body)
}
