package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToBalanceSheet(t *testing.T) {
	var buf = `
	{
		"symbol": "STOCK1",
                "annualReports": [
                     {
                         "fiscalDateEnding": "2022-12-31",
                         "reportedCurrency": "USD",
                         "totalAssets": "127243000000",
                         "totalCurrentAssets": "29118000000",
                         "cashAndCashEquivalentsAtCarryingValue": "7886000000",
                         "cashAndShortTermInvestments": "7886000000",
                         "inventory": "1552000000",
                         "currentNetReceivables": "14209000000",
                         "totalNonCurrentAssets": "96874000000",
                         "propertyPlantEquipment": "5334000000",
                         "accumulatedDepreciationAmortizationPPE": "13361000000",
                         "intangibleAssets": "67133000000",
                         "intangibleAssetsExcludingGoodwill": "11184000000",
                         "goodwill": "55949000000",
                         "investments": "None",
                         "longTermInvestments": "142000000",
                         "shortTermInvestments": "852000000",
                         "otherCurrentAssets": "2610000000",
                         "otherNonCurrentAssets": "None",
                         "totalLiabilities": "105222000000",
                         "totalCurrentLiabilities": "31505000000",
                         "currentAccountsPayable": "4051000000",
                         "deferredRevenue": "15531000000",
                         "currentDebt": "9511000000",
                         "shortTermDebt": "4760000000",
                         "totalNonCurrentLiabilities": "83414000000",
                         "capitalLeaseObligations": "164000000",
                         "longTermDebt": "47190000000",
                         "currentLongTermDebt": "4676000000",
                         "longTermDebtNoncurrent": "46189000000",
                         "shortLongTermDebtTotal": "107759000000",
                         "otherCurrentLiabilities": "9788000000",
                         "otherNonCurrentLiabilities": "12243000000",
                         "totalShareholderEquity": "21944000000",
                         "treasuryStock": "169484000000",
                         "retainedEarnings": "149825000000",
                         "commonStock": "58343000000",
                         "commonStockSharesOutstanding": "906091977"
                     }
		],
		"quarterlyReports": [
                    {
                        "fiscalDateEnding": "2023-09-30",
                        "reportedCurrency": "USD",
                        "totalAssets": "129321000000",
                        "totalCurrentAssets": "27705000000",
                        "cashAndCashEquivalentsAtCarryingValue": "7257000000",
                        "cashAndShortTermInvestments": "7257000000",
                        "inventory": "1399000000",
                        "currentNetReceivables": "6039000000",
                        "totalNonCurrentAssets": "100035000000",
                        "propertyPlantEquipment": "5369000000",
                        "accumulatedDepreciationAmortizationPPE": "12848000000",
                        "intangibleAssets": "70874000000",
                        "intangibleAssetsExcludingGoodwill": "11278000000",
                        "goodwill": "59596000000",
                        "investments": "None",
                        "longTermInvestments": "None",
                        "shortTermInvestments": "3721000000",
                        "otherCurrentAssets": "2581000000",
                        "otherNonCurrentAssets": "None",
                        "totalLiabilities": "106165000000",
                        "totalCurrentLiabilities": "30606000000",
                        "currentAccountsPayable": "3342000000",
                        "deferredRevenue": "15002000000",
                        "currentDebt": "12814000000",
                        "shortTermDebt": "6414000000",
                        "totalNonCurrentLiabilities": "75560000000",
                        "capitalLeaseObligations": "3283000000",
                        "longTermDebt": "50664000000",
                        "currentLongTermDebt": "6400000000",
                        "longTermDebtNoncurrent": "48828000000",
                        "shortLongTermDebtTotal": "84575000000",
                        "otherCurrentLiabilities": "8126000000",
                        "otherNonCurrentLiabilities": "12081000000",
                        "totalShareholderEquity": "23081000000",
                        "treasuryStock": "169640000000",
                        "retainedEarnings": "149506000000",
                        "commonStock": "59313000000",
                        "commonStockSharesOutstanding": "912800000"
			}
		]
	}
`
	balanceSheet, err := toBalanceSheet([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "STOCK1", balanceSheet.Symbol)

	assert.EqualStrings(t, "2022-12-31", balanceSheet.AnnualReports[0].FiscalDateEnding)
	assert.EqualStrings(t, "USD", balanceSheet.AnnualReports[0].ReportedCurrency)

	assert.EqualStrings(t, "2023-09-30", balanceSheet.QuarterlyReports[0].FiscalDateEnding)
	assert.EqualStrings(t, "USD", balanceSheet.QuarterlyReports[0].ReportedCurrency)
}
