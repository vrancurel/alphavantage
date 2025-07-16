package alphavantage

import (
	"encoding/json"
	"fmt"
)

// ETFProfile represents the ETF profile and holdings data.
type ETFProfile struct {
	NetAssets         string       `json:"net_assets"`
	NetExpenseRatio   string       `json:"net_expense_ratio"`
	PortfolioTurnover string       `json:"portfolio_turnover"`
	DividendYield     string       `json:"dividend_yield"`
	InceptionDate     string       `json:"inception_date"`
	Leveraged         string       `json:"leveraged"`
	Sectors           []ETF        `json:"sectors"`
	Holdings          []ETFHolding `json:"holdings"`
}

// ETF  represents sector weightings in the ETF.
type ETF struct {
	Sector string `json:"sector"`
	Weight string `json:"weight"`
}

// ETFHolding represents an individual holding in the ETF.
type ETFHolding struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Weight      string `json:"weight"`
}

// toETFProfile parses the JSON response into the ETFProfile struct.
func toETFProfile(buf []byte) (*ETFProfile, error) {
	etfProfile := &ETFProfile{}
	if err := json.Unmarshal(buf, etfProfile); err != nil {
		return nil, err
	}
	return etfProfile, nil
}

// ETFProfileData fetches and returns the ETF profile and holdings data for the specified ETF symbol.
func (c *Client) ETFProfileData(symbol string) (*ETFProfile, error) {
	const function = "ETF_PROFILE"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toETFProfile(body)
}
