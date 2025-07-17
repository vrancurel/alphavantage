package alphavantage

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Analytics represent the Metadata and the Payload
type Analytics struct {
	MetaData MetaData                   `json:"meta_data"`
	Payload  map[string]CalculationData `json:"payload"`
}

// MetaData represents the metadata in the API response.
type MetaData struct {
	Symbols  string `json:"symbols"`
	MinDt    string `json:"min_dt"`
	MaxDt    string `json:"max_dt"`
	OHLC     string `json:"ohlc"`
	Interval string `json:"interval"`
}

// CalculationData represents the payload data containing calculation results.
type CalculationData struct {
	Min              map[string]float64 `json:"MIN"`
	Max              map[string]float64 `json:"MAX"`
	Mean             map[string]float64 `json:"MEAN"`
	Median           map[string]float64 `json:"MEDIAN"`
	CumulativeReturn map[string]float64 `json:"CUMULATIVE_RETURN"`
	Variance         map[string]float64 `json:"VARIANCE"`
	StdDev           map[string]float64 `json:"STDDEV"`
	Drawdown         Drawdown           `json:"MAX_DRAWDOWN"`
	Histogram        Histogram          `json:"HISTOGRAM"`
	Autocorrelation  map[string]float64 `json:"AUTOCORRELATION"`
	Covariance       CorrelationData    `json:"COVARIANCE"`
	Correlation      CorrelationData    `json:"CORRELATION"`
}

// DrawdownRange represents the range of a drawdown.
type DrawdownRange struct {
	StartDrawdown string `json:"start_drawdown"`
	EndDrawdown   string `json:"end_drawdown"`
}

// DrawdownData represents the data for a drawdown.
type DrawdownData struct {
	MaxDrawdown   float64       `json:"max_drawdown"`
	DrawdownRange DrawdownRange `json:"drawdown_range"`
}

// Drawdown represents drawdown data for a symbol.
type Drawdown map[string]DrawdownData

// HistogramData represents the data for a histogram.
type HistogramData struct {
	BinCount []int     `json:"bin_count"`
	BinEdges []float64 `json:"bin_edges"`
}

// Histogram represents a histogram for a symbol.
type Histogram map[string]HistogramData

// CorrelationData represents the correlation data within CalculationData.
type CorrelationData struct {
	Index       []string    `json:"index"`
	Correlation [][]float64 `json:"correlation"`
}

func toAnalytics(buf []byte) (*Analytics, error) {
	analytics := &Analytics{}
	if err := json.Unmarshal(buf, analytics); err != nil {
		return nil, err
	}

	return analytics, nil
}

// Analytics fetches the analytics data for the specified symbols and
// calculations within the specified time range.  It takes a list of
// symbols, a list of calculations, a start time, an end time, an OHLC
// data type, and an interval.  The time range format depends on the
// interval.
func (c *Client) Analytics(symbols []string,
	calculations []string,
	startTime time.Time,
	endTime time.Time,
	ohlc AnalyticsOhlc,
	interval AnalyticsInterval) (*Analytics, error) {
	const functionName = "analytics"
	_symbols := strings.Join(symbols, ",")
	var _rangeStart string
	var _rangeEnd string
	if interval == AnalyticsIntervalDaily ||
		interval == AnalyticsIntervalWeekly ||
		interval == AnalyticsIntervalMonthly {
		_rangeStart = startTime.Format("2006-01-02")
		_rangeEnd = endTime.Format("2006-01-02")
	} else {
		_rangeStart = startTime.Format("2006-01-02T15:04:05")
		_rangeEnd = endTime.Format("2006-01-02T15:04:05")
	}
	_ohlc := string(ohlc)
	_calculations := strings.Join(calculations, ",")
	url := fmt.Sprintf("%s/timeseries/%s?SYMBOLS=%s&CALCULATIONS=%s&RANGE=%s&RANGE=%s&OHLC=%s&INTERVAL=%s&apikey=%s",
		baseURLApi, functionName, _symbols, _calculations, _rangeStart, _rangeEnd, _ohlc, interval, c.apiKey)

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	analytics, err := toAnalytics(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return analytics, nil
}

// Analytics2 fetches the analytics data for the specified symbols and
// calculations using a range unit.  It takes a list of symbols, a
// list of calculations, a range factor, a range unit, an OHLC data
// type, and an interval.
func (c *Client) Analytics2(symbols []string,
	calculations []string,
	rangeFactor int,
	rangeUnit AnalyticsRangeUnit,
	ohlc AnalyticsOhlc,
	interval AnalyticsInterval) (*Analytics, error) {
	const functionName = "analytics"
	_symbols := strings.Join(symbols, ",")
	var _range string
	if rangeUnit == AnalyticsRangeUnitFull {
		_range = string(rangeUnit)
	} else {
		_range = strconv.Itoa(rangeFactor) + string(rangeUnit)
	}
	_ohlc := string(ohlc)
	_calculations := strings.Join(calculations, ",")
	url := fmt.Sprintf("%s/timeseries/%s?SYMBOLS=%s&CALCULATIONS=%s&RANGE=%s&OHLC=%s&INTERVAL=%s&apikey=%s",
		baseURLApi, functionName, _symbols, _calculations, _range, _ohlc, interval, c.apiKey)

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	analytics, err := toAnalytics(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return analytics, nil
}
