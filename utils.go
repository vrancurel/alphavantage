package alphavantage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// PrettyPrint formats a given data structure into indented JSON.
func PrettyPrint(s interface{}) ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json: %v", err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

// PrevQuarter calculates the previous quarter for a given year and quarter.
func PrevQuarter(year int, quarter int) (int, int, error) {
	if quarter < 0 || quarter > 3 {
		return -1, -1, errors.New("Quarter must be between 0 and 3")
	}
	prevYear := year
	prevQuarter := quarter - 1
	if prevQuarter == -1 {
		prevQuarter = 3
		prevYear--
	}
	return prevYear, prevQuarter, nil
}

// GetYearAndQuarter extracts the year and quarter from a date string.
func GetYearAndQuarter(dateString string) (int, int, error) {
	reportDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return -1, -1, fmt.Errorf("Parse error: %v", err)
	}
	// Calculate the quarter
	quarter := (int(reportDate.Month()) - 1) / 3
	return reportDate.Year(), quarter, nil
}

// GetPrevYearAndQuarter calculates the previous year and quarter for a given date string.
func GetPrevYearAndQuarter(dateString string) (int, int, error) {
	reportDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return -1, -1, fmt.Errorf("Parse error: %v", err)
	}
	// Calculate the quarter
	quarter := (int(reportDate.Month()) - 1) / 3
	return PrevQuarter(reportDate.Year(), quarter)
}

// GetPrevYear calculates the previous year and quarter for a given date string.
func GetPrevYear(dateString string) (int, int, error) {
	reportDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return -1, -1, fmt.Errorf("Parse error: %v", err)
	}
	// Calculate the quarter
	quarter := (int(reportDate.Month()) - 1) / 3
	return reportDate.Year() - 1, quarter, nil
}

// GetCashFlowMostRecentQuarterlyReport retrieves the most recent quarterly report from a CashFlow struct.
func GetCashFlowMostRecentQuarterlyReport(cf *CashFlow) (*CfQuarterlyReport, error) {
	var mostRecent CfQuarterlyReport
	mostRecentDate := time.Time{}
	for _, report := range cf.QuarterlyReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetCashFlowMostRecentAnnualReport retrieves the most recent annual report from a CashFlow struct.
func GetCashFlowMostRecentAnnualReport(cf *CashFlow) (*CfAnnualReport, error) {
	var mostRecent CfAnnualReport
	mostRecentDate := time.Time{}
	for _, report := range cf.AnnualReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetCashFlowHistoricalQuarterlyReport retrieves a specific quarterly report from a CashFlow struct.
func GetCashFlowHistoricalQuarterlyReport(cf *CashFlow, year int, quarter int) (*CfQuarterlyReport, error) {
	if quarter < 0 || quarter > 3 {
		return nil, errors.New("Quarter must be between 0 and 3")
	}
	targetDates := []string{
		fmt.Sprintf("%d-03-31", year),
		fmt.Sprintf("%d-06-30", year),
		fmt.Sprintf("%d-09-30", year),
		fmt.Sprintf("%d-12-31", year),
	}
	for _, report := range cf.QuarterlyReports {
		if report.FiscalDateEnding == targetDates[quarter] {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetCashFlowHistoricalAnnualReport retrieves a specific annual report from a CashFlow struct.
func GetCashFlowHistoricalAnnualReport(cf *CashFlow, year int) (*CfAnnualReport, error) {
	for _, report := range cf.AnnualReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.Year() == year {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetBalanceSheetMostRecentQuarterlyReport retrieves the most recent quarterly report from a BalanceSheet struct.
func GetBalanceSheetMostRecentQuarterlyReport(bs *BalanceSheet) (*BsQuarterlyReport, error) {
	var mostRecent BsQuarterlyReport
	mostRecentDate := time.Time{}
	for _, report := range bs.QuarterlyReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetBalanceSheetMostRecentAnnualReport retrieves the most recent annual report from a BalanceSheet struct.
func GetBalanceSheetMostRecentAnnualReport(bs *BalanceSheet) (*BsAnnualReport, error) {
	var mostRecent BsAnnualReport
	mostRecentDate := time.Time{}
	for _, report := range bs.AnnualReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetBalanceSheetHistoricalQuarterlyReport retrieves a specific quarterly report from a BalanceSheet struct.
func GetBalanceSheetHistoricalQuarterlyReport(bs *BalanceSheet, year int, quarter int) (*BsQuarterlyReport, error) {
	if quarter < 0 || quarter > 3 {
		return nil, errors.New("Quarter must be between 0 and 3")
	}
	targetDates := []string{
		fmt.Sprintf("%d-03-31", year),
		fmt.Sprintf("%d-06-30", year),
		fmt.Sprintf("%d-09-30", year),
		fmt.Sprintf("%d-12-31", year),
	}
	for _, report := range bs.QuarterlyReports {
		if report.FiscalDateEnding == targetDates[quarter] {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetBalanceSheetHistoricalAnnualReport retrieves a specific annual report from a BalanceSheet struct.
func GetBalanceSheetHistoricalAnnualReport(bs *BalanceSheet, year int) (*BsAnnualReport, error) {
	for _, report := range bs.AnnualReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.Year() == year {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetIncomeStatementMostRecentQuarterlyReport retrieves the most recent quarterly report from an IncomeStatement struct.
func GetIncomeStatementMostRecentQuarterlyReport(is *IncomeStatement) (*IsQuarterlyReport, error) {
	var mostRecent IsQuarterlyReport
	mostRecentDate := time.Time{}
	for _, report := range is.QuarterlyReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetIncomeStatementMostRecentAnnualReport retrieves the most recent annual report from an IncomeStatement struct.
func GetIncomeStatementMostRecentAnnualReport(is *IncomeStatement) (*IsAnnualReport, error) {
	var mostRecent IsAnnualReport
	mostRecentDate := time.Time{}
	for _, report := range is.AnnualReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetIncomeStatementHistoricalQuarterlyReport retrieves a specific quarterly report from an IncomeStatement struct.
func GetIncomeStatementHistoricalQuarterlyReport(is *IncomeStatement, year int, quarter int) (*IsQuarterlyReport, error) {
	if quarter < 0 || quarter > 3 {
		return nil, errors.New("Quarter must be between 0 and 3")
	}
	targetDates := []string{
		fmt.Sprintf("%d-03-31", year),
		fmt.Sprintf("%d-06-30", year),
		fmt.Sprintf("%d-09-30", year),
		fmt.Sprintf("%d-12-31", year),
	}
	for _, report := range is.QuarterlyReports {
		if report.FiscalDateEnding == targetDates[quarter] {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetIncomeStatementHistoricalAnnualReport retrieves a specific annual report from an IncomeStatement struct.
func GetIncomeStatementHistoricalAnnualReport(is *IncomeStatement, year int) (*IsAnnualReport, error) {
	for _, report := range is.AnnualReports {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.Year() == year {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetEarningsMostRecentQuarterlyReport retrieves the most recent quarterly report from an Earnings struct.
func GetEarningsMostRecentQuarterlyReport(e *Earnings) (*QuarterlyEarnings, error) {
	var mostRecent QuarterlyEarnings
	mostRecentDate := time.Time{}
	for _, report := range e.QuarterlyEarnings {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetEarningsMostRecentAnnualReport retrieves the most recent annual report from an Earnings struct.
func GetEarningsMostRecentAnnualReport(e *Earnings) (*AnnualEarnings, error) {
	var mostRecent AnnualEarnings
	mostRecentDate := time.Time{}
	for _, report := range e.AnnualEarnings {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.After(mostRecentDate) {
			mostRecent = report
			mostRecentDate = reportDate
		}
	}
	return &mostRecent, nil
}

// GetEarningsHistoricalQuarterlyReport retrieves a specific quarterly report from an Earnings struct.
func GetEarningsHistoricalQuarterlyReport(e *Earnings, year int, quarter int) (*QuarterlyEarnings, error) {
	if quarter < 0 || quarter > 3 {
		return nil, errors.New("Quarter must be between 0 and 3")
	}
	targetDates := []string{
		fmt.Sprintf("%d-03-31", year),
		fmt.Sprintf("%d-06-30", year),
		fmt.Sprintf("%d-09-30", year),
		fmt.Sprintf("%d-12-31", year),
	}
	for _, report := range e.QuarterlyEarnings {
		if report.FiscalDateEnding == targetDates[quarter] {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetEarningsHistoricalAnnualReport retrieves a specific annual report from an Earnings struct.
func GetEarningsHistoricalAnnualReport(e *Earnings, year int) (*AnnualEarnings, error) {
	for _, report := range e.AnnualEarnings {
		reportDate, err := time.Parse("2006-01-02", report.FiscalDateEnding)
		if err != nil {
			return nil, fmt.Errorf("Parse error: %v", err)
		}
		if reportDate.Year() == year {
			return &report, nil
		}
	}
	return nil, errors.New("Not found")
}

// GetTimeSeriesMonthlyData retrieves the monthly data from a TimeSeries struct for a specific year and quarter.
func GetTimeSeriesMonthlyData(ts *TimeSeries, year int, quarter int) (*TimeSeriesData, string, error) {
	if quarter < 0 || quarter > 3 {
		return nil, "", errors.New("Quarter must be between 0 and 3")
	}
	month := (quarter + 1) * 3
	for dateStr, data := range ts.TimeSeriesMonthly {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, "", fmt.Errorf("FindTimeSeriesDataByDate: %v", err)
		}
		if date.Year() == year && int(date.Month()) == month {
			return &data, dateStr, nil
		}
	}
	return nil, "", errors.New("Not found")
}

// TranslateStringToInterval translates a string into an av.Interval
func TranslateStringToInterval(intervalStr string) (Interval, error) {
	switch intervalStr {
	case "1min":
		return Interval1Min, nil
	case "5mn":
		return Interval5Min, nil
	case "15mn":
		return Interval15Min, nil
	case "30mn":
		return Interval30Min, nil
	case "60mn":
		return Interval60Min, nil
	case "daily":
		return IntervalDaily, nil
	case "weekly":
		return IntervalWeekly, nil
	case "monthly":
		return IntervalMonthly, nil
	default:
		return "", errors.New("invalid interval string")
	}
}

// IntervalToExpirationDelay transforms an Interval to a time.Duration
func IntervalToExpirationDelay(interval Interval) time.Duration {
	switch interval {
	case Interval1Min:
		return time.Minute * 1
	case Interval5Min:
		return time.Minute * 5
	case Interval15Min:
		return time.Minute * 15
	case Interval30Min:
		return time.Minute * 30
	case Interval60Min:
		return time.Minute * 60
	case IntervalDaily:
		return time.Hour * 24
	case IntervalWeekly:
		return time.Hour * 24 * 7
	case IntervalMonthly:
		return time.Hour * 24 * 30
	default:
		panic("No such interval")
	}
}

// TranslateStringToTimeSeriesInterval translates a string into an av.TimeSeriesInterval
func TranslateStringToTimeSeriesInterval(intervalStr string) (TimeSeriesInterval, error) {
	switch intervalStr {
	case "daily":
		return TimeSeriesDaily, nil
	case "weekly":
		return TimeSeriesWeekly, nil
	case "monthly":
		return TimeSeriesMonthly, nil
	default:
		return "", errors.New("invalid interval string")
	}
}

// TimeSeriesIntervalToExpirationDelay transforms an Interval to a time.Duration
func TimeSeriesIntervalToExpirationDelay(interval TimeSeriesInterval) time.Duration {
	switch interval {
	case TimeSeriesDaily:
		return time.Hour * 24
	case TimeSeriesWeekly:
		return time.Hour * 24 * 7
	case TimeSeriesMonthly:
		return time.Hour * 24 * 30
	default:
		panic("No such time series interval")
	}
}

// TranslateStringToTimeSeriesIntervalAdjusted translates a string into an av.TimeSeriesIntervalAdjusted
func TranslateStringToTimeSeriesIntervalAdjusted(intervalStr string) (TimeSeriesIntervalAdjusted, error) {
	switch intervalStr {
	case "daily":
		return TimeSeriesDailyAdjusted, nil
	case "weekly":
		return TimeSeriesWeeklyAdjusted, nil
	case "monthly":
		return TimeSeriesMonthlyAdjusted, nil
	default:
		return "", errors.New("invalid interval string")
	}
}

// TimeSeriesIntervalAdjustedToExpirationDelay transforms an Interval to a time.Duration
func TimeSeriesIntervalAdjustedToExpirationDelay(interval TimeSeriesIntervalAdjusted) time.Duration {
	switch interval {
	case TimeSeriesDailyAdjusted:
		return time.Hour * 24
	case TimeSeriesWeeklyAdjusted:
		return time.Hour * 24 * 7
	case TimeSeriesMonthlyAdjusted:
		return time.Hour * 24 * 30
	default:
		panic("No such time series interval adjusted")
	}
}

// TranslateStringToOutputSize translates a string into an av.TimeSeriesIntervalAdjusted
func TranslateStringToOutputSize(outputSizeStr string) (OutputSize, error) {
	switch outputSizeStr {
	case "compact":
		return OutputSizeCompact, nil
	case "full":
		return OutputSizeFull, nil
	default:
		return "", errors.New("invalid output size string")
	}
}
