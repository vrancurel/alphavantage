package alphavantage

import (
	"fmt"
	"log"
	"time"
)

// Horizon represents the time period for earnings calendar lookup
type Horizon string

const (
	// ThreeMonth represents 3-month horizon
	ThreeMonth Horizon = "3month"
	// SixMonth represents 6-month horizon
	SixMonth Horizon = "6month"
	// TwelveMonth represents 12-month horizon
	TwelveMonth Horizon = "12month"
)

// EarningsEvent represents a single earnings calendar event
type EarningsEvent struct {
	Symbol           string
	Name             string
	ReportDate       time.Time
	FiscalDateEnding time.Time
	Estimate         float64
	Currency         string
}

// EarningsCalendar represents the collection of earnings events
type EarningsCalendar struct {
	Events []EarningsEvent
}

func toEarningsCalendar(data [][]string) (*EarningsCalendar, error) {
	var ec EarningsCalendar
	ec.Events = make([]EarningsEvent, 0)

	// Load Eastern Time location
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("Error loading location: %v", err)
	}

	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		if len(row) != 6 {
			log.Printf("Invalid number of fields: %v", len(row))
			continue
		}

		var event EarningsEvent
		event.Symbol = row[0]
		event.Name = row[1]

		// Parse report date
		reportDate, err := time.ParseInLocation("2006-01-02", row[2], est)
		if err != nil {
			log.Printf("Error parsing report date %v", row[2])
			continue
		}
		event.ReportDate = reportDate

		// Parse fiscal date ending
		fiscalDate, err := time.ParseInLocation("2006-01-02", row[3], est)
		if err != nil {
			log.Printf("Error parsing fiscal date %v", row[3])
			continue
		}
		event.FiscalDateEnding = fiscalDate

		// Parse estimate (if provided)
		if row[4] != "" {
			_, err := fmt.Sscanf(row[4], "%f", &event.Estimate)
			if err != nil {
				log.Printf("Error parsing estimate %v", row[4])
				continue
			}
		}

		event.Currency = row[5]
		ec.Events = append(ec.Events, event)
	}

	return &ec, nil
}

// EarningsCalendar fetches the earnings calendar data for the specified parameters
func (c *Client) EarningsCalendar(symbol string, horizon Horizon) (*EarningsCalendar, error) {
	const function = "EARNINGS_CALENDAR"

	url := fmt.Sprintf("%s/query?function=%s&apikey=%s", baseURL, function, c.apiKey)

	if symbol != "" {
		url += fmt.Sprintf("&symbol=%s", symbol)
	}

	if horizon != "" {
		url += fmt.Sprintf("&horizon=%s", horizon)
	}

	data, err := c.makeHTTPRequestForCsv(url)
	if err != nil {
		return nil, err
	}

	return toEarningsCalendar(data)
}
