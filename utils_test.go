package alphavantage

import (
	"github.com/AMekss/assert"
	"testing"
)

func TestPrevQuarter(t *testing.T) {
	year := 2019
	quarter := 1
	year, quarter, err := PrevQuarter(year, quarter)
	assert.NoError(t.Fatalf, err)
	assert.EqualInt(t, 2019, year)
	assert.EqualInt(t, 0, quarter)
	year, quarter, err = PrevQuarter(year, quarter)
	assert.NoError(t.Fatalf, err)
	assert.EqualInt(t, 2018, year)
	assert.EqualInt(t, 3, quarter)
}

func TestGetYearAndQuarter(t *testing.T) {
	year, quarter, err := GetYearAndQuarter("2019-06-30")
	assert.NoError(t.Fatalf, err)
	assert.EqualInt(t, 2019, year)
	assert.EqualInt(t, 1, quarter)
}

func TestGetPrevYearAndQuarter(t *testing.T) {
	year, quarter, err := GetPrevYearAndQuarter("2019-06-30")
	assert.NoError(t.Fatalf, err)
	assert.EqualInt(t, 2019, year)
	assert.EqualInt(t, 0, quarter)
	year, quarter, err = GetPrevYearAndQuarter("2019-03-31")
	assert.NoError(t.Fatalf, err)
	assert.EqualInt(t, 2018, year)
	assert.EqualInt(t, 3, quarter)
}

func TestGetPrevYear(t *testing.T) {
	year, quarter, err := GetPrevYear("2019-06-30")
	assert.NoError(t.Fatalf, err)
	assert.EqualInt(t, 2018, year)
	assert.EqualInt(t, 1, quarter)
	year, quarter, err = GetPrevYear("2019-03-31")
	assert.NoError(t.Fatalf, err)
	assert.EqualInt(t, 2018, year)
	assert.EqualInt(t, 0, quarter)
}
