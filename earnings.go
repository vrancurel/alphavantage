package alphavantage

import (
	"encoding/json"
	"fmt"
)

// AnnualEarnings represents the annual earnings data for a company.
type AnnualEarnings struct {
	FiscalDateEnding string    `json:"fiscalDateEnding"`
	ReportedEPS      AVFloat64 `json:"reportedEPS"`
}

// QuarterlyEarnings represents the quarterly earnings data for a company.
type QuarterlyEarnings struct {
	FiscalDateEnding   string    `json:"fiscalDateEnding"`
	ReportedDate       string    `json:"reportedDate"`
	ReportedEPS        AVFloat64 `json:"reportedEPS"`
	EstimatedEPS       AVFloat64 `json:"estimatedEPS"`
	Surprise           AVFloat64 `json:"surprise"`
	SurprisePercentage AVFloat64 `json:"surprisePercentage"`
}

// Earnings represents the earnings data for a company, including both annual and quarterly earnings.
type Earnings struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}

func toEarnings(buf []byte) (*Earnings, error) {
	earnings := &Earnings{}
	if err := json.Unmarshal(buf, earnings); err != nil {
		return nil, err
	}
	return earnings, nil
}

// Earnings fetches and returns the earnings data for the specified company symbol.
func (c *Client) Earnings(symbol string) (*Earnings, error) {
	const function = "EARNINGS"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toEarnings(body)
}
