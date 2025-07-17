package alphavantage

import (
	"github.com/AMekss/assert"
	"testing"
)

func TestToCompanyOverview(t *testing.T) {
	var buf = `
{
    "Symbol": "STOCK1",
    "AssetType": "Common Stock",
    "Name": "Stock1 Company",
    "Description": "Stock1 company builds blah.",
    "CIK": "523423433434",
    "Exchange": "NYSE",
    "Currency": "USD",
    "Country": "USA",
    "Sector": "TECHNOLOGY",
    "Industry": "COMPUTER & OFFICE EQUIPMENT",
    "Address": "XXX",
    "FiscalYearEnd": "December",
    "LatestQuarter": "2023-03-31",
    "MarketCapitalization": "112577577000",
    "EBITDA": "12644000000",
    "PERatio": "55.33",
    "PEGRatio": "1.276",
    "BookValue": "23.79",
    "DividendPerShare": "6.6",
    "DividendYield": "0.0544",
    "EPS": "2.22",
    "RevenuePerShareTTM": "66.97",
    "ProfitMargin": "0.0303",
    "OperatingMarginTTM": "0.132",
    "ReturnOnAssetsTTM": "0.0376",
    "ReturnOnEquityTTM": "0.101",
    "RevenueTTM": "60585001000",
    "GrossProfitTTM": "32688000000",
    "DilutedEPSTTM": "2.22",
    "QuarterlyEarningsGrowthYOY": "0.253",
    "QuarterlyRevenueGrowthYOY": "0.004",
    "AnalystTargetPrice": "140.79",
    "TrailingPE": "55.33",
    "ForwardPE": "15.55",
    "PriceToSalesRatioTTM": "2.108",
    "PriceToBookRatio": "6.75",
    "EVToRevenue": "2.969",
    "EVToEBITDA": "25.81",
    "Beta": "0.851",
    "52WeekHigh": "149.31",
    "52WeekLow": "111.29",
    "50DayMovingAverage": "126.63",
    "200DayMovingAverage": "133.09",
    "SharesOutstanding": "908045000",
    "DividendDate": "2023-06-10",
    "ExDividendDate": "2023-05-09"
}
`
	var companyOverview, err = toCompanyOverview([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "STOCK1", companyOverview.Symbol)
	assert.EqualStrings(t, "Common Stock", companyOverview.AssetType)
	assert.EqualStrings(t, "Stock1 Company", companyOverview.Name)
	assert.EqualStrings(t, "523423433434", companyOverview.CIK)
	assert.EqualStrings(t, "NYSE", companyOverview.Exchange)
	assert.EqualStrings(t, "USD", companyOverview.Currency)
	assert.EqualInt(t, 112577577000, int(companyOverview.MarketCapitalization.Value))
	assert.EqualInt(t, 12644000000, int(companyOverview.EBITDA.Value))
	assert.EqualFloat64(t, 55.33, companyOverview.PERatio.Value)
	assert.EqualFloat64(t, 1.276, companyOverview.PEGRatio.Value)
	assert.EqualFloat64(t, 23.79, companyOverview.BookValue.Value)
	assert.EqualFloat64(t, 6.6, companyOverview.DividendPerShare.Value)
	assert.EqualFloat64(t, 0.0544, companyOverview.DividendYield.Value)
	assert.EqualFloat64(t, 2.22, companyOverview.EPS.Value)
	assert.EqualFloat64(t, 66.97, companyOverview.RevenuePerShareTTM.Value)
	assert.EqualFloat64(t, 0.0303, companyOverview.ProfitMargin.Value)
	assert.EqualFloat64(t, 0.132, companyOverview.OperatingMarginTTM.Value)
	assert.EqualFloat64(t, 0.0376, companyOverview.ReturnOnAssetsTTM.Value)
	assert.EqualFloat64(t, 0.101, companyOverview.ReturnOnEquityTTM.Value)
	assert.EqualInt(t, 60585001000, int(companyOverview.RevenueTTM.Value))
	assert.EqualInt(t, 32688000000, int(companyOverview.GrossProfitTTM.Value))
	assert.EqualFloat64(t, 2.22, companyOverview.DilutedEPSTTM.Value)
	assert.EqualFloat64(t, 0.253, companyOverview.QuarterlyEarningsGrowthYOY.Value)
	assert.EqualFloat64(t, 0.004, companyOverview.QuarterlyRevenueGrowthYOY.Value)
	assert.EqualFloat64(t, 140.79, companyOverview.AnalystTargetPrice.Value)
	assert.EqualFloat64(t, 55.33, companyOverview.TrailingPE.Value)
	assert.EqualFloat64(t, 15.55, companyOverview.ForwardPE.Value)
	assert.EqualFloat64(t, 2.108, companyOverview.PriceToSalesRatioTTM.Value)
	assert.EqualFloat64(t, 6.75, companyOverview.PriceToBookRatio.Value)
	assert.EqualFloat64(t, 2.969, companyOverview.EVToRevenue.Value)
	assert.EqualFloat64(t, 25.81, companyOverview.EVToEBITDA.Value)
	assert.EqualFloat64(t, 0.851, companyOverview.Beta.Value)
	assert.EqualFloat64(t, 149.31, companyOverview.Week52High.Value)
	assert.EqualFloat64(t, 111.29, companyOverview.Week52Low.Value)
	assert.EqualFloat64(t, 126.63, companyOverview.MovingAverage50Day.Value)
	assert.EqualFloat64(t, 133.09, companyOverview.MovingAverage200Day.Value)
	assert.EqualInt(t, 908045000, int(companyOverview.SharesOutstanding.Value))
}
