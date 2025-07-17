package alphavantage

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Interval represents the possible interval types
type Interval string

// TimeSeriesInterval is the type for time series data
type TimeSeriesInterval string

// TimeSeriesIntervalAdjusted is the type for adjusted time series data
type TimeSeriesIntervalAdjusted string

// SortType represents the type for sorting data.
type SortType string

// OutputSize is the type for data output
type OutputSize string

// AnalyticsCalculation represents the type for analytics calculations.
type AnalyticsCalculation string

// AnalyticsRangeUnit represents the range unit for analytics data.
type AnalyticsRangeUnit string

// AnalyticsOhlc represents the OHLC (Open, High, Low, Close) data type.
type AnalyticsOhlc string

// AnalyticsInterval represents the interval for analytics data.
type AnalyticsInterval string

// SeriesType type of series for indicators
type SeriesType string

// AVFloat64 represents a custom float type to handle "None" value.
type AVFloat64 struct {
	Value float64
}

// UnmarshalJSON custom unmarshaller to handle "None" value.
func (cf *AVFloat64) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case float64:
		cf.Value = v
	case string:
		if v == "None" || v == "-" {
			cf.Value = 0.0
		} else {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("unexpected string value: %s: %v", v, err)
			}
			cf.Value = f
		}
	default:
		return fmt.Errorf("unexpected value type: %T", v)
	}
	return nil
}

// MarshalJSON method to handle custom JSON marshaling
func (cf AVFloat64) MarshalJSON() ([]byte, error) {
	if cf.Value == 0 {
		return []byte("\"None\""), nil
	}
	return json.Marshal(strconv.FormatFloat(cf.Value, 'f', -1, 64))
}

// AVPercent represents a custom float type to handle percentage values.
type AVPercent struct {
	Value float64
}

// UnmarshalJSON custom unmarshaller to handle "%" value.
func (cf *AVPercent) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case float64:
		cf.Value = v
	case string:
		if v == "None" || v == "-" {
			cf.Value = 0.0
		} else {
			_v := strings.ReplaceAll(v, "%", "")
			f, err := strconv.ParseFloat(_v, 64)
			if err != nil {
				return fmt.Errorf("unexpected string value: %s: %v", v, err)
			}
			cf.Value = f
		}
	default:
		return fmt.Errorf("unexpected value type: %T", v)
	}
	return nil
}

// MarshalJSON method to handle custom JSON marshaling
func (cf AVPercent) MarshalJSON() ([]byte, error) {
	if cf.Value == 0 {
		return []byte("\"None\""), nil
	}
	return json.Marshal(strconv.FormatFloat(cf.Value, 'f', 4, 64) + "%")
}

// AVInt represents a custom int type to handle "None" value.
type AVInt struct {
	Value int
}

// UnmarshalJSON custom unmarshaller to handle "None" value.
func (cf *AVInt) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case int:
		cf.Value = v
	case string:
		if v == "None" || v == "-" {
			cf.Value = 0
		} else {
			i, err := strconv.Atoi(v)
			if err != nil {
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("unexpected string value: %s: %v", v, err)
				}
				cf.Value = int(f)
			} else {
				cf.Value = i
			}
		}
	default:
		return fmt.Errorf("unexpected value type: %T", v)
	}
	return nil
}

// MarshalJSON method to handle custom JSON marshaling
func (cf AVInt) MarshalJSON() ([]byte, error) {
	if cf.Value == 0 {
		return []byte("\"None\""), nil
	}
	return json.Marshal(strconv.Itoa(cf.Value))
}

// MarshalCSV custom marshaller to handle CSV marshaling.
func (cf AVInt) MarshalCSV() ([]byte, error) {
	if cf.Value == 0 {
		return []byte("\"None\""), nil
	}
	return strconv.AppendInt(nil, int64(cf.Value), 10), nil
}

const (
	// Interval1Min represents the 1 minute interval.
	Interval1Min = Interval("1min")
	// Interval5Min represents the 5 minute interval.
	Interval5Min = Interval("5mn")
	// Interval15Min represents the 15 minute interval.
	Interval15Min = Interval("15mn")
	// Interval30Min represents the 30 minute interval.
	Interval30Min = Interval("30mn")
	// Interval60Min represents the 60 minute interval.
	Interval60Min = Interval("60mn")
	// IntervalDaily represents the daily interval.
	IntervalDaily = Interval("daily")
	// IntervalWeekly represents the weekly interval.
	IntervalWeekly = Interval("weekly")
	// IntervalMonthly represents the monthly interval.
	IntervalMonthly = Interval("monthly")

	// DateFormat is the date format used by the API
	DateFormat = "2006-01-02"
	// DateTimeFormat datetime format used by the API
	DateTimeFormat = "2006-01-02 15:04:05"

	// TimeSeriesDaily - Time series for daily prices
	TimeSeriesDaily = TimeSeriesInterval("TIME_SERIES_DAILY")
	// TimeSeriesWeekly - Time series for weekly prices
	TimeSeriesWeekly = TimeSeriesInterval("TIME_SERIES_WEEKLY")
	// TimeSeriesMonthly - Time series for monthly prices
	TimeSeriesMonthly = TimeSeriesInterval("TIME_SERIES_MONTHLY")

	// TimeSeriesDailyAdjusted represents daily adjusted time series data.
	TimeSeriesDailyAdjusted = TimeSeriesIntervalAdjusted("TIME_SERIES_DAILY_ADJUSTED")
	// TimeSeriesWeeklyAdjusted represents weekly adjusted time series data.
	TimeSeriesWeeklyAdjusted = TimeSeriesIntervalAdjusted("TIME_SERIES_WEEKLY_ADJUSTED")
	// TimeSeriesMonthlyAdjusted represents monthly adjusted time series data.
	TimeSeriesMonthlyAdjusted = TimeSeriesIntervalAdjusted("TIME_SERIES_MONTHLY_ADJUSTED")

	// OutputSizeCompact is for the latest 100 items
	OutputSizeCompact = OutputSize("compact")
	// OutputSizeFull is for the full-length time series
	OutputSizeFull = OutputSize("full")

	// SortTypeLatest represents sorting by latest.
	SortTypeLatest = SortType("LATEST")
	// SortTypeEarliest represents sorting by earliest.
	SortTypeEarliest = SortType("EARLIEST")
	// SortTypeRelevance represents sorting by relevance.
	SortTypeRelevance = SortType("RELEVANCE")

	// AnalyticsCalculationMin represents the minimum return calculation.
	AnalyticsCalculationMin = AnalyticsCalculation("MIN")
	// AnalyticsCalculationMax represents the maximum return calculation.
	AnalyticsCalculationMax = AnalyticsCalculation("MAX")
	// AnalyticsCalculationMean represents the mean return calculation.
	AnalyticsCalculationMean = AnalyticsCalculation("MEAN")
	// AnalyticsCalculationMedian represents the median return calculation.
	AnalyticsCalculationMedian = AnalyticsCalculation("MEDIAN")
	// AnalyticsCalculationCumulativeReturn represents the cumulative return calculation.
	AnalyticsCalculationCumulativeReturn = AnalyticsCalculation("CUMULATIVE_RETURN")
	// AnalyticsCalculationVariance represents the variance calculation.
	AnalyticsCalculationVariance = AnalyticsCalculation("VARIANCE")
	// AnalyticsCalculationStdDev represents the standard deviation calculation.
	AnalyticsCalculationStdDev = AnalyticsCalculation("STDDEV")
	// AnalyticsCalculationMaxDrawdown represents the maximum drawdown calculation.
	AnalyticsCalculationMaxDrawdown = AnalyticsCalculation("MAX_DRAWDOWN")
	// AnalyticsCalculationHistogram represents the histogram calculation.
	AnalyticsCalculationHistogram = AnalyticsCalculation("HISTOGRAM")
	// AnalyticsCalculationAutocorrelation represents the autocorrelation calculation.
	AnalyticsCalculationAutocorrelation = AnalyticsCalculation("AUTOCORRELATION")
	// AnalyticsCalculationCovariance represents the covariance calculation.
	AnalyticsCalculationCovariance = AnalyticsCalculation("COVARIANCE")
	// AnalyticsCalculationCorrelation represents the correlation calculation.
	AnalyticsCalculationCorrelation = AnalyticsCalculation("CORRELATION")

	// AnalyticsRangeUnitFull represents the full range unit.
	AnalyticsRangeUnitFull = AnalyticsRangeUnit("full")
	// AnalyticsRangeUnitDay represents the day range unit.
	AnalyticsRangeUnitDay = AnalyticsRangeUnit("day")
	// AnalyticsRangeUnitWeek represents the week range unit.
	AnalyticsRangeUnitWeek = AnalyticsRangeUnit("week")
	// AnalyticsRangeUnitMonth represents the month range unit.
	AnalyticsRangeUnitMonth = AnalyticsRangeUnit("month")
	// AnalyticsRangeUnitYear represents the year range unit.
	AnalyticsRangeUnitYear = AnalyticsRangeUnit("year")
	// AnalyticsRangeUnitMinute represents the minute range unit.
	AnalyticsRangeUnitMinute = AnalyticsRangeUnit("minute")
	// AnalyticsRangeUnitHour represents the hour range unit.
	AnalyticsRangeUnitHour = AnalyticsRangeUnit("hour")

	// AnalyticsOhlcOpen represents the open OHLC data type.
	AnalyticsOhlcOpen = AnalyticsOhlc("open")
	// AnalyticsOhlcHigh represents the high OHLC data type.
	AnalyticsOhlcHigh = AnalyticsOhlc("high")
	// AnalyticsOhlcLow represents the low OHLC data type.
	AnalyticsOhlcLow = AnalyticsOhlc("low")
	// AnalyticsOhlcClose represents the close OHLC data type.
	AnalyticsOhlcClose = AnalyticsOhlc("close")

	// AnalyticsInterval1Min represents the 1 minute interval.
	AnalyticsInterval1Min = AnalyticsInterval("1min")
	// AnalyticsInterval5Min represents the 5 minute interval.
	AnalyticsInterval5Min = AnalyticsInterval("5mn")
	// AnalyticsInterval15Min represents the 15 minute interval.
	AnalyticsInterval15Min = AnalyticsInterval("15mn")
	// AnalyticsInterval30Min represents the 30 minute interval.
	AnalyticsInterval30Min = AnalyticsInterval("30mn")
	// AnalyticsInterval60Min represents the 60 minute interval.
	AnalyticsInterval60Min = AnalyticsInterval("60mn")
	// AnalyticsIntervalDaily represents the daily interval.
	AnalyticsIntervalDaily = AnalyticsInterval("DAILY")
	// AnalyticsIntervalWeekly represents the weekly interval.
	AnalyticsIntervalWeekly = AnalyticsInterval("WEEKLY")
	// AnalyticsIntervalMonthly represents the monthly interval.
	AnalyticsIntervalMonthly = AnalyticsInterval("MONTHLY")

	// SeriesTypeOpen - Series Type Open
	SeriesTypeOpen = SeriesType("open")
	// SeriesTypeHigh - Series Type High
	SeriesTypeHigh = SeriesType("high")
	// SeriesTypeLow - Series Type Low
	SeriesTypeLow = SeriesType("low")
	// SeriesTypeClose - Series Type Close
	SeriesTypeClose = SeriesType("close")
)
