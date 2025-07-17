package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToInsiderTransactions(t *testing.T) {
	// Sample JSON response
	jsonData := []byte(`{
		"data": [
			{
				"transaction_date": "2025-03-31",
				"ticker": "STOCK1",
				"executive": "Executive1",
				"executive_title": "Director",
				"security_type": "Promised Fee Share",
				"acquisition_or_disposal": "A",
				"shares": "393.0",
				"share_price": "248.66"
			},
			{
				"transaction_date": "2025-03-31",
				"ticker": "STOCK1",
				"executive": "Executive2",
				"executive_title": "Director",
				"security_type": "Promised Fee Share",
				"acquisition_or_disposal": "A",
				"shares": "367.0",
				"share_price": "248.66"
			}
		]
	}`)

	// Parse the JSON
	result, err := toInsiderTransactions(jsonData)

	// Verify parsing was successful
	assert.NoError(t, err)
	assert.EqualInt(t, len(result.Data), 2)

	// Verify first transaction was parsed correctly
	transaction := result.Data[0]
	assert.EqualStrings(t, "2025-03-31", transaction.TransactionDate)
	assert.EqualStrings(t, "STOCK1", transaction.Ticker)
	assert.EqualStrings(t, "Executive1", transaction.Executive)
	assert.EqualStrings(t, "Director", transaction.ExecutiveTitle)
	assert.EqualStrings(t, "Promised Fee Share", transaction.SecurityType)
	assert.EqualStrings(t, "A", transaction.AcquisitionOrDisposal)
	assert.EqualStrings(t, "393.0", transaction.Shares)
	assert.EqualStrings(t, "248.66", transaction.SharePrice)

	// Verify second transaction was parsed correctly
	transaction = result.Data[1]
	assert.EqualStrings(t, "2025-03-31", transaction.TransactionDate)
	assert.EqualStrings(t, "STOCK1", transaction.Ticker)
	assert.EqualStrings(t, "Executive2", transaction.Executive)
	assert.EqualStrings(t, "Director", transaction.ExecutiveTitle)
	assert.EqualStrings(t, "Promised Fee Share", transaction.SecurityType)
	assert.EqualStrings(t, "A", transaction.AcquisitionOrDisposal)
	assert.EqualStrings(t, "367.0", transaction.Shares)
	assert.EqualStrings(t, "248.66", transaction.SharePrice)
}
