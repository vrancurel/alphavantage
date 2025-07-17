package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToListingStatus(t *testing.T) {
	var buf = [][]string{
		[]string{"symbol", "name", "exchange", "assetType", "ipoDate", "delistingDate", "status"},
		[]string{"STOCK1", "Stock1 Inc", "NYSE", "Stock", "1999-11-18", "null", "Active"},
		[]string{"STOCK2", "Stock2 Corp", "NYSE", "Stock", "2016-10-18", "null", "Active"},
	}

	listingStatus, err := toListingStatus(buf, false)
	assert.NoError(t.Fatalf, err)

	assert.EqualInt(t, 2, len(listingStatus.SymbolStatuses))
}
