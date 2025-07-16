package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToEarnings(t *testing.T) {
	var buf = `
{
        "symbol": "STOCK1",
        "annualEarnings": [
            {
                "fiscalDateEnding": "2023-12-31",
                "reportedEPS": "9.61"
            },
            {
                "fiscalDateEnding": "2022-12-31",
                "reportedEPS": "9.12"
            }
        ],
        "quarterlyEarnings": [
            {
                "fiscalDateEnding": "2023-12-31",
                "reportedDate": "2024-01-24",
                "reportedEPS": "3.87",
                "estimatedEPS": "3.78",
                "surprise": "0.09",
                "surprisePercentage": "2.381"
            }
        ]
    }
`
	earnings, err := toEarnings([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "STOCK1", earnings.Symbol)

	assert.EqualStrings(t, "2023-12-31", earnings.AnnualEarnings[0].FiscalDateEnding)
	assert.EqualStrings(t, "2023-12-31", earnings.QuarterlyEarnings[0].FiscalDateEnding)
}
