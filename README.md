# alphavantage

**Unofficial** Go/Golang www.alphavantage.co API implementation

**Disclaimer**: This library is not associated with alphavantage or any of its affiliates or subsidiaries. If you use this library, you should contact them to make sure they are okay with how you intend to use it. Use this lib at your own risk.

API doc reference: https://www.alphavantage.co/documentation/

**Note**: Requests are synchronised and throttled automatically to not flood the API servers. One request every 15 seconds is possible at the moment.

## Usage

### New client

```go
package main

import (
	"github.com/sklinkert/alphavantage"
	log "github.com/sirupsen/logrus" // optional
)

func main() {
	avClient := alphavantage.New("MYAPIKEY")
 	// ...
}
```

### TimeSeries (prices)

```go
series, err := avClient.TimeSeries("TICKER1", alphavantage.TimeSeriesDaily, alphavantage.OutPutSizeCompact)
if err != nil {
	log.WithError(err).Fatal("TimeSeries() failed")
}
for date, price := range series.TimeSeriesDaily {
	log.Infof("%s: Open=%f High=%f Low=%f Close=%f Volume=%d", date, price.Open, price.High, price.Low, price.Close, price.Volume)
}
```

### Indicator STOCH

```go
indicators, err := avClient.IndicatorStoch("EURUSD", alphavantage.IntervalDaily)
if err != nil {
	log.WithError(err).Fatal("IndicatorStoch() failed")
}

// Loop over all indicators
for date, indicator := range indicators.TechnicalAnalysis {
	log.Infof("%s: SlowK=%f SlowD=%f", date, indicator.SlowK, indicator.SlowD)
}

// Get the most recent one
latestDate, latest := indicators.Latest()
log.Infof("Latest: %s: SlowK=%f SlowD=%f", latestDate, latest.SlowK, latest.SlowD)

// Get today only
today := indicators.Today()
log.Infof("Today: SlowK=%f SlowD=%f", today.SlowK, today.SlowD)

// By specific date
indicator := indicators.ByDate(time.Now())
if indicator != nil {
	log.Infof("SlowK=%f SlowD=%f", indicator.SlowK, indicator.SlowD)
}

```

### Indicator SMA

```go
period := 200
indicators, err := avClient.IndicatorSMA("EURUSD", alphavantage.IntervalDaily, period, alphavantage.SeriesTypeClose)
if err != nil {
	log.WithError(err).Fatal("IndicatorSMA() failed")
}

// See more examples at Indicator STOCH section
```

### Global Quote

```go
globalQuote, err := avClient.GlobalQuote("EURUSD")
if err != nil {
	log.WithError(err).Fatal("GlobalQuote() failed")
}
log.Infof("%s %f", globalQuote.LatestTradingDay, globalQuote.Price)
```

### Analytics

```go
// with date range
analytics, err := avClient.Analytics([]string{"TICKER1", "TICKER2"}, "2023-03-03", "2024-03-01", alphavantage.IntervalDaily)
if err != nil {
	log.WithError(err).Fatal("Analytics() failed")
}
log.Infof("%v", analytics)

// without date range
analytics, err = avClient.Analytics2([]string{"TICKER1", "TICKER2"}, alphavantage.IntervalDaily)
if err != nil {
	log.WithError(err).Fatal("Analytics() failed")
}
log.Infof("%v", analytics)
```

### Balance Sheet

```go
balanceSheet, err := avClient.BalanceSheet("TICKER")
if err != nil {
	log.WithError(err).Fatal("BalanceSheet() failed")
}
log.Infof("%v", balanceSheet)
```

### Cash Flow

```go
cashFlow, err := avClient.CashFlow("TICKER")
if err != nil {
	log.WithError(err).Fatal("CashFlow() failed")
}
log.Infof("%v", cashFlow)
```

### Company Overview

```go
companyOverview, err := avClient.CompanyOverview("TICKER")
if err != nil {
	log.WithError(err).Fatal("CompanyOverview() failed")
}
log.Infof("%v", companyOverview)
```

### Earnings Calendar

```go
// with 3month horizon
earningsCalendar, err := avClient.EarningsCalendar("TICKER", alphavantage.ThreeMonth)
if err != nil {
	log.WithError(err).Fatal("EarningsCalendar() failed")
}
log.Infof("%v", earningsCalendar)

// with 6month horizon
earningsCalendar, err = avClient.EarningsCalendar("TICKER", alphavantage.SixMonth)
if err != nil {
	log.WithError(err).Fatal("EarningsCalendar() failed")
}
log.Infof("%v", earningsCalendar)

// with 12month horizon
earningsCalendar, err = avClient.EarningsCalendar("TICKER", alphavantage.TwelveMonth)
if err != nil {
	log.WithError(err).Fatal("EarningsCalendar() failed")
}
log.Infof("%v", earningsCalendar)
```

### Earnings

```go
earnings, err := avClient.Earnings("TICKER")
if err != nil {
	log.WithError(err).Fatal("Earnings() failed")
}
log.Infof("%v", earnings)
```

### Earnings Call Transcript

```go
transcript, err := avClient.EarningsCallTranscript("TICKER", "2022-03-31")
if err != nil {
	log.WithError(err).Fatal("EarningsCallTranscript() failed")
}
log.Infof("%v", transcript)
```

### ETF Profile

```go
etfProfile, err := avClient.ETFProfileData("VOO")
if err != nil {
	log.WithError(err).Fatal("ETFProfileData() failed")
}
log.Infof("%v", etfProfile)
```

### Historical Options

```go
// with specific date
date := time.Date(2022, 11, 18, 0, 0, 0, 0, time.UTC)
options, err := avClient.HistoricalOptions("TICKER2", &date)
if err != nil {
	log.WithError(err).Fatal("HistoricalOptions() failed")
}
log.Infof("%v", options)

// without specific date
options, err = avClient.HistoricalOptions("TICKER2", nil)
if err != nil {
	log.WithError(err).Fatal("HistoricalOptions() failed")
}
log.Infof("%v", options)
```

### Income Statement

```go
incomeStatement, err := avClient.IncomeStatement("TICKER")
if err != nil {
	log.WithError(err).Fatal("IncomeStatement() failed")
}
log.Infof("%v", incomeStatement)
```

### Indicator EMA

```go
indicatorEMA, err := avClient.IndicatorEMA("EURUSD", alphavantage.IntervalDaily, 200, alphavantage.SeriesTypeClose)
if err != nil {
	log.WithError(err).Fatal("IndicatorEMA() failed")
}
log.Infof("%v", indicatorEMA)
```

### Insider Transactions

```go
insiderTransactions, err := avClient.InsiderTransactions("TICKER")
if err != nil {
	log.WithError(err).Fatal("InsiderTransactions() failed")
}
log.Infof("%v", insiderTransactions)
```

### Listing Status

```go
// for delisted stocks
listingStatus, err := avClient.ListingStatus(true)
if err != nil {
	log.WithError(err).Fatal("ListingStatus() failed")
}
log.Infof("%v", listingStatus)

// for active stocks
listingStatus, err = avClient.ListingStatus(false)
if err != nil {
	log.WithError(err).Fatal("ListingStatus() failed")
}
log.Infof("%v", listingStatus)
```

### News Sentiment

```go
newsSentiment, err := avClient.NewsSentiment("TICKER2", alphavantage.SortLatest, 10, "", "")
if err != nil {
	log.WithError(err).Fatal("NewsSentiment() failed")
}
log.Infof("%v", newsSentiment)
```

### Time Series Adjusted

```go
// with compact output size
timeSeriesAdjusted, err := avClient.TimeSeriesAdjusted("TICKER", alphavantage.TimeSeriesDaily, alphavantage.OutPutSizeCompact)
if err != nil {
	log.WithError(err).Fatal("TimeSeriesAdjusted() failed")
}
log.Infof("%v", timeSeriesAdjusted)

// with full output size
timeSeriesAdjusted, err = avClient.TimeSeriesAdjusted("TICKER", alphavantage.TimeSeriesDaily, alphavantage.OutPutSizeFull)
if err != nil {
	log.WithError(err).Fatal("TimeSeriesAdjusted() failed")
}
log.Infof("%v", timeSeriesAdjusted)
```
