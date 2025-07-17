package alphavantage

import (
	"encoding/json"
	"fmt"
)

// CashFlow represents the cash flow data for a company.
type CashFlow struct {
	Symbol           string              `json:"symbol"`
	AnnualReports    []CfAnnualReport    `json:"annualReports"`
	QuarterlyReports []CfQuarterlyReport `json:"quarterlyReports"`
}

// CfAnnualReport represents an annual report in the cash flow.
type CfAnnualReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         AVInt  `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            AVInt  `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           AVInt  `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              AVInt  `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   AVInt  `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      AVInt  `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       AVInt  `json:"capitalExpenditures"`
	ChangeInReceivables                                       AVInt  `json:"changeInReceivables"`
	ChangeInInventory                                         AVInt  `json:"changeInInventory"`
	ProfitLoss                                                AVInt  `json:"profitLoss"`
	CashflowFromInvestment                                    AVInt  `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     AVInt  `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     AVInt  `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        AVInt  `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             AVInt  `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     AVInt  `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            AVInt  `json:"dividendPayout"`
	DividendPayoutCommonStock                                 AVInt  `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              AVInt  `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         AVInt  `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet AVInt  `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      AVInt  `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            AVInt  `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           AVInt  `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            AVInt  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      AVInt  `json:"changeInExchangeRate"`
	NetIncome                                                 AVInt  `json:"netIncome"`
}

// CfQuarterlyReport represents a quarterly report in the cash flow.
type CfQuarterlyReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         AVInt  `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            AVInt  `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           AVInt  `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              AVInt  `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   AVInt  `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      AVInt  `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       AVInt  `json:"capitalExpenditures"`
	ChangeInReceivables                                       AVInt  `json:"changeInReceivables"`
	ChangeInInventory                                         AVInt  `json:"changeInInventory"`
	ProfitLoss                                                AVInt  `json:"profitLoss"`
	CashflowFromInvestment                                    AVInt  `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     AVInt  `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     AVInt  `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        AVInt  `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             AVInt  `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     AVInt  `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            AVInt  `json:"dividendPayout"`
	DividendPayoutCommonStock                                 AVInt  `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              AVInt  `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         AVInt  `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet AVInt  `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      AVInt  `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            AVInt  `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           AVInt  `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            AVInt  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      AVInt  `json:"changeInExchangeRate"`
	NetIncome                                                 AVInt  `json:"netIncome"`
}

func toCashFlow(buf []byte) (*CashFlow, error) {
	cashFlow := &CashFlow{}
	if err := json.Unmarshal(buf, cashFlow); err != nil {
		return nil, err
	}
	return cashFlow, nil
}

// CashFlow fetches and returns the cash flow data for the specified company symbol.
func (c *Client) CashFlow(symbol string) (*CashFlow, error) {
	const function = "CASH_FLOW"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toCashFlow(body)
}
