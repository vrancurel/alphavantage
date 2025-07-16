package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToCashFlow(t *testing.T) {
	var buf = `
	{
		"symbol": "STOCK1",
                "annualReports": [
                     {
            "fiscalDateEnding": "2022-12-31",
            "reportedCurrency": "USD",
            "operatingCashflow": "10435000000",
            "paymentsForOperatingActivities": "2430000000",
            "proceedsFromOperatingActivities": "None",
            "changeInOperatingLiabilities": "213000000",
            "changeInOperatingAssets": "468000000",
            "depreciationDepletionAndAmortization": "4802000000",
            "capitalExpenditures": "1346000000",
            "changeInReceivables": "539000000",
            "changeInInventory": "-71000000",
            "profitLoss": "1639000000",
            "cashflowFromInvestment": "-4202000000",
            "cashflowFromFinancing": "-4958000000",
            "proceedsFromRepaymentsOfShortTermDebt": "217000000",
            "paymentsForRepurchaseOfCommonStock": "None",
            "paymentsForRepurchaseOfEquity": "None",
            "paymentsForRepurchaseOfPreferredStock": "None",
            "dividendPayout": "5948000000",
            "dividendPayoutCommonStock": "5948000000",
            "dividendPayoutPreferredStock": "None",
            "proceedsFromIssuanceOfCommonStock": "None",
            "proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet": "7804000000",
            "proceedsFromIssuanceOfPreferredStock": "None",
            "proceedsFromRepurchaseOfEquity": "-407000000",
            "proceedsFromSaleOfTreasuryStock": "None",
            "changeInCashAndCashEquivalents": "None",
            "changeInExchangeRate": "None",
            "netIncome": "1639000000"
                     }
		],
		"quarterlyReports": [
                    {
            "fiscalDateEnding": "2023-09-30",
            "reportedCurrency": "USD",
            "operatingCashflow": "3056000000",
            "paymentsForOperatingActivities": "None",
            "proceedsFromOperatingActivities": "None",
            "changeInOperatingLiabilities": "None",
            "changeInOperatingAssets": "None",
            "depreciationDepletionAndAmortization": "1093000000",
            "capitalExpenditures": "281000000",
            "changeInReceivables": "None",
            "changeInInventory": "None",
            "profitLoss": "1714000000",
            "cashflowFromInvestment": "-1953000000",
            "cashflowFromFinancing": "-3132000000",
            "proceedsFromRepaymentsOfShortTermDebt": "9000000",
            "paymentsForRepurchaseOfCommonStock": "0",
            "paymentsForRepurchaseOfEquity": "0",
            "paymentsForRepurchaseOfPreferredStock": "None",
            "dividendPayout": "1515000000",
            "dividendPayoutCommonStock": "1515000000",
            "dividendPayoutPreferredStock": "None",
            "proceedsFromIssuanceOfCommonStock": "None",
            "proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet": "154000000",
            "proceedsFromIssuanceOfPreferredStock": "None",
            "proceedsFromRepurchaseOfEquity": "0",
            "proceedsFromSaleOfTreasuryStock": "None",
            "changeInCashAndCashEquivalents": "None",
            "changeInExchangeRate": "None",
            "netIncome": "1704000000"
  	            }
		]
	}
`
	cashFlow, err := toCashFlow([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "STOCK1", cashFlow.Symbol)

	assert.EqualStrings(t, "2022-12-31", cashFlow.AnnualReports[0].FiscalDateEnding)
	assert.EqualStrings(t, "USD", cashFlow.AnnualReports[0].ReportedCurrency)
	assert.EqualInt(t, 10435000000, int(cashFlow.AnnualReports[0].OperatingCashflow.Value))

	assert.EqualStrings(t, "2023-09-30", cashFlow.QuarterlyReports[0].FiscalDateEnding)
	assert.EqualStrings(t, "USD", cashFlow.QuarterlyReports[0].ReportedCurrency)
}
