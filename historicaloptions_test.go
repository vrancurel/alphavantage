package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToHistoricalOptionsData(t *testing.T) {
	var buf = `
{
    "symbol": "STOCK1",
    "endpoint": "Historical Options",
    "message": "success",
    "data": [
        {
            "contractID": "STOCK1240906C00095000",
            "symbol": "STOCK1",
            "expiration": "2024-09-06",
            "strike": "95.00",
            "type": "call",
            "last": "0.00",
            "mark": "105.93",
            "bid": "104.20",
            "bid_size": "37",
            "ask": "107.65",
            "ask_size": "41",
            "volume": "0",
            "open_interest": "0",
            "date": "2024-09-06",
            "implied_volatility": "6.28020",
            "delta": "0.99267",
            "gamma": "0.00031",
            "theta": "-2.49365",
            "vega": "0.00213",
            "rho": "0.00256"
        },
        {
            "contractID": "STOCK1240906P00095000",
            "symbol": "STOCK1",
            "expiration": "2024-09-06",
            "strike": "95.00",
            "type": "put",
            "last": "0.00",
            "mark": "0.01",
            "bid": "0.00",
            "bid_size": "0",
            "ask": "0.06",
            "ask_size": "50",
            "volume": "0",
            "open_interest": "20",
            "date": "2024-09-06",
            "implied_volatility": "4.65829",
            "delta": "-0.00071",
            "gamma": "0.00005",
            "theta": "-0.21922",
            "vega": "0.00026",
            "rho": "-0.00000"
        }
    ]
}
`
	optionsData, err := toHistoricalOptionsData([]byte(buf))
	assert.NoError(t.Fatalf, err)

	// Asserting symbol and general information
	assert.EqualStrings(t, "STOCK1", optionsData.Symbol)
	assert.EqualStrings(t, "Historical Options", optionsData.Endpoint)
	assert.EqualStrings(t, "success", optionsData.Message)

	// Asserting the first contract details (Call option)
	firstContract := optionsData.Data[0]
	assert.EqualStrings(t, "STOCK1240906C00095000", firstContract.ContractID)
	assert.EqualStrings(t, "call", firstContract.Type)
	assert.EqualStrings(t, "2024-09-06", firstContract.Expiration)
	assert.EqualFloat64(t, 95.00, firstContract.Strike.Value)
	assert.EqualFloat64(t, 0.99267, firstContract.Delta.Value)

	// Asserting the second contract details (Put option)
	secondContract := optionsData.Data[1]
	assert.EqualStrings(t, "STOCK1240906P00095000", secondContract.ContractID)
	assert.EqualStrings(t, "put", secondContract.Type)
	assert.EqualFloat64(t, -0.00071, secondContract.Delta.Value)
}
