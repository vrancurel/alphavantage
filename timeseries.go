package alphavantage

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// TimeSeries represents the overall struct for time series
type TimeSeries struct {
	Metadata          TimeSeriesMetadata        `json:"Meta Data"`
	TimeSeriesDaily   map[string]TimeSeriesData `json:"Time Series (Daily)"`
	TimeSeriesWeekly  map[string]TimeSeriesData `json:"Weekly Time Series"`
	TimeSeriesMonthly map[string]TimeSeriesData `json:"Monthly Time Series"`
}

// TimeSeriesMetadata is the metadata subset of TimeSeries
type TimeSeriesMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

// TimeSeriesData is a subset of TimeSeries
type TimeSeriesData struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume uint64  `json:"5. volume,string"`
}

// TimeSeriesAdjusted - like TimeSeries, but inclused dividends and adjusted close
type TimeSeriesAdjusted struct {
	Metadata          TimeSeriesMetadata                `json:"Meta Data"`
	TimeSeriesDaily   map[string]TimeSeriesAdjustedData `json:"Time Series (Daily)"`
	TimeSeriesWeekly  map[string]TimeSeriesAdjustedData `json:"Weekly Adjusted Time Series"`
	TimeSeriesMonthly map[string]TimeSeriesAdjustedData `json:"Monthly Adjusted Time Series"`
}

// TimeSeriesAdjustedData - like TimeSeries, but inclused dividends and adjusted close
type TimeSeriesAdjustedData struct {
	Open             float64 `json:"1. open,string"`
	High             float64 `json:"2. high,string"`
	Low              float64 `json:"3. low,string"`
	Close            float64 `json:"4. close,string"`
	AdjustedClose    float64 `json:"5. adjusted close,string"`
	Volume           uint64  `json:"6. volume,string"`
	DividendAmount   float64 `json:"7. dividend amount,string"`
	SplitCoefficient float64 `json:"8. split coefficient,string"`
}

func toTimeSeries(buf []byte) (*TimeSeries, error) {
	timeSeries := &TimeSeries{}
	if err := json.Unmarshal(buf, timeSeries); err != nil {
		return nil, err
	}
	return timeSeries, nil
}

func toTimeSeriesAdjusted(buf []byte) (*TimeSeriesAdjusted, error) {
	timeSeries := &TimeSeriesAdjusted{}
	if err := json.Unmarshal(buf, timeSeries); err != nil {
		return nil, err
	}
	return timeSeries, nil
}

// TimeSeriesAdjusted fetches the time series for given symbol from API.
// The order of dates in returned object is random because it's a map.
func (c *Client) TimeSeriesAdjusted(symbol string, interval TimeSeriesIntervalAdjusted, outputSize OutputSize) (*TimeSeriesAdjusted, error) {
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s&outputsize=%s", baseURL, interval, symbol, c.apiKey, outputSize)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	timeSeries, err := toTimeSeriesAdjusted(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return timeSeries, nil
}

// TimeSeries fetches the time series for given symbol from API.
// The order of dates in returned object is random because it's a map.
func (c *Client) TimeSeries(symbol string, interval TimeSeriesInterval, outputSize OutputSize) (*TimeSeries, error) {
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s&outputsize=%s", baseURL, interval, symbol, c.apiKey, outputSize)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	timeSeries, err := toTimeSeries(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return timeSeries, nil
}

// getFilledData returns the data subset for the filled interval
func (ts *TimeSeries) getFilledData() map[string]TimeSeriesData {
	if len(ts.TimeSeriesDaily) > 0 {
		return ts.TimeSeriesDaily
	} else if len(ts.TimeSeriesWeekly) > 0 {
		return ts.TimeSeriesWeekly
	}
	return nil
}

// Len returns the number of data items
func (ts *TimeSeries) Len() int {
	fd := ts.getFilledData()
	return len(fd)
}

// Latest returns the most recent item
func (ts *TimeSeries) Latest() (date string, latest *TimeSeriesData) {
	datasets := ts.getFilledData()
	dates := make([]string, len(datasets))
	for date := range datasets {
		dates = append(dates, date)
	}
	sort.Strings(dates)
	date = dates[len(dates)-1]
	latestVal, _ := datasets[date]
	latest = &latestVal
	return
}

// Today returns dataset for today.
func (ts *TimeSeries) Today() *TimeSeriesData {
	today := time.Now()
	return ts.ByDate(today)
}

// ByDate returns the dataset for the given date.
func (ts *TimeSeries) ByDate(date time.Time) *TimeSeriesData {
	dataset := ts.getFilledData()
	day := date.Format(DateFormat)
	item, exists := dataset[day]
	if !exists {
		return nil
	}
	return &item
}
