package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToETFProfile(t *testing.T) {
	var buf = `
{
    "net_assets": "323800000000",
    "net_expense_ratio": "0.002",
    "portfolio_turnover": "0.08",
    "dividend_yield": "0.0055",
    "inception_date": "1999-03-10",
    "leveraged": "NO",
    "sectors": [
        {
            "sector": "INFORMATION TECHNOLOGY",
            "weight": "0.489"
        }
    ],
    "holdings": [
        {
            "symbol": "STOCK1",
            "description": "STOCK1 INC",
            "weight": "0.0907"
        }
    ]
}
`
	etfProfile, err := toETFProfile([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "323800000000", etfProfile.NetAssets)
	assert.EqualStrings(t, "0.002", etfProfile.NetExpenseRatio)
	assert.EqualStrings(t, "STOCK1", etfProfile.Holdings[0].Symbol)
	assert.EqualStrings(t, "INFORMATION TECHNOLOGY", etfProfile.Sectors[0].Sector)

}
