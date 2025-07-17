package alphavantage

import (
	"encoding/json"
	"fmt"
)

// IndicatorEMA represents the overall struct for stochastics indicator
// Example https://www.alphavantage.co/query?function=STOCH&symbol=MSFT&interval=daily&apikey=demo
type IndicatorEMA struct {
	Metadata          IndicatorEMAMetadata            `json:"Meta Data"`
	TechnicalAnalysis map[string]TechnicalEMAAnalysis `json:"Technical Analysis: EMA"`
}

// IndicatorEMAMetadata is the metadata subset of IndicatorEMA
type IndicatorEMAMetadata struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	TimePeriod    int    `json:"5: Time Period"`
	SeriesType    string `json:"6: Series Type"`
	TimeZone      string `json:"7: Time Zone"`
}

// TechnicalEMAAnalysis is the EMA indicator subset of IndicatorEMA
type TechnicalEMAAnalysis struct {
	EMA float64 `json:",string"`
}

func toIndicatorEMA(buf []byte) (*IndicatorEMA, error) {
	indicatorEMA := &IndicatorEMA{}
	if err := json.Unmarshal(buf, indicatorEMA); err != nil {
		return nil, err
	}
	return indicatorEMA, nil
}

// IndicatorEMA fetches the "EMA" indicators for given symbol from API.
// The order of dates in TechnicalAnalysis is random because it's a map.
func (c *Client) IndicatorEMA(symbol string, interval Interval, timePeriod int, seriesType SeriesType) (*IndicatorEMA, error) {
	const functionName = "EMA"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&interval=%s&time_period=%d&series_type=%s&apikey=%s",
		baseURL, functionName, symbol, interval, timePeriod, seriesType, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	indicator, err := toIndicatorEMA(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return indicator, nil
}
