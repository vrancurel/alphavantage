package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToIndicatorEMA(t *testing.T) {
	var buf = `
	{
		"Meta Data": {
			"1: Symbol": "USDEUR",
			"2: Indicator": "Simple Moving Average (EMA)",
			"3: Last Refreshed": "2020-04-30",
			"4: Interval": "weekly",
			"5: Time Period": 10,
			"6: Series Type": "open",
			"7: Time Zone": "US/Eastern"
		},
		"Technical Analysis: EMA": {
			"2020-04-30": {
				"EMA": "0.9120"
			},
			"2020-04-24": {
				"EMA": "0.9118"
			},
			"2020-04-17": {
				"EMA": "0.9111"
			},
			"2020-04-10": {
				"EMA": "0.9098"
			},
			"2020-04-03": {
				"EMA": "0.9080"
			}
		}
    }
`
	indicator, err := toIndicatorEMA([]byte(buf))
	assert.NoError(t, err)
	assert.EqualStrings(t, "USDEUR", indicator.Metadata.Symbol)
	assert.EqualInt(t, 5, len(indicator.TechnicalAnalysis))

	ta1, exists := indicator.TechnicalAnalysis["2020-04-24"]
	if !exists {
		assert.Panic(t, "entry for 2020-04-24 is missing")
	}
	assert.EqualFloat64(t, 0.9118, ta1.EMA)
}
