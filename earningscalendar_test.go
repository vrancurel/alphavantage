package alphavantage

import (
	"testing"
	"time"

	"github.com/AMekss/assert"
)

var timeEST = mustLoadLocation("America/New_York")

func mustLoadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err) // Panic if the location cannot be loaded
	}
	return loc
}

func TestToEarningsCalendar(t *testing.T) {
	var buf = [][]string{
		{"symbol", "name", "reportDate", "fiscalDateEnding", "estimate", "currency"},
		{"STOCK1", "Stock1 Company", "2025-01-23", "2024-12-31", "0.39", "USD"},
		{"STOCK2", "Stock2 Company", "2025-03-12", "2024-12-31", "", "USD"},
	}

	calendar, err := toEarningsCalendar(buf)
	assert.NoError(t.Fatalf, err)

	assert.EqualInt(t, 2, len(calendar.Events))

	// Test first event
	event := calendar.Events[0]
	assert.EqualStrings(t, "STOCK1", event.Symbol)
	assert.EqualStrings(t, "Stock1 Company", event.Name)

	expectedReportDate := time.Date(2025, 1, 23, 0, 0, 0, 0, timeEST)
	assert.True(t, event.ReportDate.Equal(expectedReportDate))

	expectedFiscalDate := time.Date(2024, 12, 31, 0, 0, 0, 0, timeEST)
	assert.True(t, event.FiscalDateEnding.Equal(expectedFiscalDate))

	assert.EqualFloat64(t, 0.39, event.Estimate)
	assert.EqualStrings(t, "USD", event.Currency)

	// Test second event (with empty estimate)
	event = calendar.Events[1]
	assert.EqualStrings(t, "STOCK2", event.Symbol)
	assert.EqualFloat64(t, 0.0, event.Estimate)
}
