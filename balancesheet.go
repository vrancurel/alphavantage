package alphavantage

import (
	"encoding/json"
	"fmt"
)

// BalanceSheet represents the balance sheet data for a company.
type BalanceSheet struct {
	Symbol           string              `json:"symbol"`
	AnnualReports    []BsAnnualReport    `json:"annualReports"`
	QuarterlyReports []BsQuarterlyReport `json:"quarterlyReports"`
}

// BsAnnualReport represents an annual report in the balance sheet.
type BsAnnualReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            AVInt  `json:"totalAssets"`
	TotalCurrentAssets                     AVInt  `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  AVInt  `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            AVInt  `json:"cashAndShortTermInvestments"`
	Inventory                              AVInt  `json:"inventory"`
	CurrentNetReceivables                  AVInt  `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  AVInt  `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 AVInt  `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPpe AVInt  `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       AVInt  `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      AVInt  `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               AVInt  `json:"goodwill"`
	Investments                            AVInt  `json:"investments"`
	LongTermInvestments                    AVInt  `json:"longTermInvestments"`
	ShortTermInvestments                   AVInt  `json:"shortTermInvestments"`
	OtherCurrentAssets                     AVInt  `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  AVInt  `json:"otherNonCurrentAssets"`
	TotalLiabilities                       AVInt  `json:"totalLiabilities"`
	TotalCurrentLiabilities                AVInt  `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 AVInt  `json:"currentAccountsPayable"`
	DeferredRevenue                        AVInt  `json:"deferredRevenue"`
	CurrentDebt                            AVInt  `json:"currentDebt"`
	ShortTermDebt                          AVInt  `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             AVInt  `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                AVInt  `json:"capitalLeaseObligations"`
	LongTermDebt                           AVInt  `json:"longTermDebt"`
	CurrentLongTermDebt                    AVInt  `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 AVInt  `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 AVInt  `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                AVInt  `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             AVInt  `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 AVInt  `json:"totalShareholderEquity"`
	TreasuryStock                          AVInt  `json:"treasuryStock"`
	RetainedEarnings                       AVInt  `json:"retainedEarnings"`
	CommonStock                            AVInt  `json:"commonStock"`
	CommonStockSharesOutstanding           AVInt  `json:"commonStockSharesOutstanding"`
}

// BsQuarterlyReport represents a quarterly report in the balance sheet.
type BsQuarterlyReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            AVInt  `json:"totalAssets"`
	TotalCurrentAssets                     AVInt  `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  AVInt  `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            AVInt  `json:"cashAndShortTermInvestments"`
	Inventory                              AVInt  `json:"inventory"`
	CurrentNetReceivables                  AVInt  `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  AVInt  `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 AVInt  `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPpe AVInt  `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       AVInt  `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      AVInt  `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               AVInt  `json:"goodwill"`
	Investments                            AVInt  `json:"investments"`
	LongTermInvestments                    AVInt  `json:"longTermInvestments"`
	ShortTermInvestments                   AVInt  `json:"shortTermInvestments"`
	OtherCurrentAssets                     AVInt  `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  AVInt  `json:"otherNonCurrentAssets"`
	TotalLiabilities                       AVInt  `json:"totalLiabilities"`
	TotalCurrentLiabilities                AVInt  `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 AVInt  `json:"currentAccountsPayable"`
	DeferredRevenue                        AVInt  `json:"deferredRevenue"`
	CurrentDebt                            AVInt  `json:"currentDebt"`
	ShortTermDebt                          AVInt  `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             AVInt  `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                AVInt  `json:"capitalLeaseObligations"`
	LongTermDebt                           AVInt  `json:"longTermDebt"`
	CurrentLongTermDebt                    AVInt  `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 AVInt  `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 AVInt  `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                AVInt  `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             AVInt  `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 AVInt  `json:"totalShareholderEquity"`
	TreasuryStock                          AVInt  `json:"treasuryStock"`
	RetainedEarnings                       AVInt  `json:"retainedEarnings"`
	CommonStock                            AVInt  `json:"commonStock"`
	CommonStockSharesOutstanding           AVInt  `json:"commonStockSharesOutstanding"`
}

func toBalanceSheet(buf []byte) (*BalanceSheet, error) {
	balanceSheet := &BalanceSheet{}
	if err := json.Unmarshal(buf, balanceSheet); err != nil {
		return nil, err
	}
	return balanceSheet, nil
}

// BalanceSheet fetches and returns the balance sheet data for the specified company symbol.
func (c *Client) BalanceSheet(symbol string) (*BalanceSheet, error) {
	const function = "BALANCE_SHEET"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toBalanceSheet(body)
}
