package alphavantage

import (
	"encoding/json"
	"fmt"
)

// CompanyOverview represents the company overview data for a company.
type CompanyOverview struct {
	Symbol                     string    `json:"Symbol"`
	AssetType                  string    `json:"AssetType"`
	Name                       string    `json:"Name"`
	Description                string    `json:"Description"`
	CIK                        string    `json:"CIK"`
	Exchange                   string    `json:"Exchange"`
	Currency                   string    `json:"Currency"`
	Country                    string    `json:"Country"`
	Sector                     string    `json:"Sector"`
	Industry                   string    `json:"Industry"`
	Address                    string    `json:"Address"`
	FiscalYearEnd              string    `json:"FiscalYearEnd"`
	LatestQuarter              string    `json:"LatestQuarter"` // "2023-03-31"
	MarketCapitalization       AVInt     `json:"MarketCapitalization,string"`
	EBITDA                     AVInt     `json:"EBITDA,string"`
	PERatio                    AVFloat64 `json:"PERatio,string"`
	PEGRatio                   AVFloat64 `json:"PEGRatio,string"`
	BookValue                  AVFloat64 `json:"BookValue,string"`
	DividendPerShare           AVFloat64 `json:"DividendPerShare,string"`
	DividendYield              AVFloat64 `json:"DividendYield,string"`
	EPS                        AVFloat64 `json:"EPS,string"`
	RevenuePerShareTTM         AVFloat64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               AVFloat64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         AVFloat64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          AVFloat64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          AVFloat64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 AVInt     `json:"RevenueTTM,string"`
	GrossProfitTTM             AVInt     `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              AVFloat64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY AVFloat64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  AVFloat64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         AVFloat64 `json:"AnalystTargetPrice,string"`
	AnalystRatingStrongBuy     AVInt     `json:"AnalystRatingStrongBuy,string"`
	AnalystRatingBuy           AVInt     `json:"AnalystRatingBuy,string"`
	AnalystRatingHold          AVInt     `json:"AnalystRatingHold,string"`
	AnalystRatingSell          AVInt     `json:"AnalystRatingSell,string"`
	AnalystRatingStrongSell    AVInt     `json:"AnalystRatingStrongSell,string"`
	TrailingPE                 AVFloat64 `json:"TrailingPE,string"`
	ForwardPE                  AVFloat64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       AVFloat64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           AVFloat64 `json:"PriceToBookRatio,string"`
	EVToRevenue                AVFloat64 `json:"EVToRevenue,string"`
	EVToEBITDA                 AVFloat64 `json:"EVToEBITDA,string"`
	Beta                       AVFloat64 `json:"Beta,string"`
	Week52High                 AVFloat64 `json:"52WeekHigh,string"`
	Week52Low                  AVFloat64 `json:"52WeekLow,string"`
	MovingAverage50Day         AVFloat64 `json:"50DayMovingAverage,string"`
	MovingAverage200Day        AVFloat64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          AVInt     `json:"SharesOutstanding,string"`
	DividendDate               string    `json:"DividendDate"`   //  "2023-06-10"
	ExDividendDate             string    `json:"ExDividendDate"` // "2023-05-09"
}

func toCompanyOverview(buf []byte) (*CompanyOverview, error) {
	var overview CompanyOverview
	err := json.Unmarshal(buf, &overview)
	if err != nil {
		return nil, err
	}
	return &overview, nil
}

// CompanyOverview fetches and returns the company overview data for the specified company symbol.
func (c *Client) CompanyOverview(symbol string) (*CompanyOverview, error) {
	const function = "OVERVIEW"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	return toCompanyOverview(body)
}
