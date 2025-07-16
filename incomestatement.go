package alphavantage

import (
	"encoding/json"
	"fmt"
)

// IncomeStatement represents the income statement data for a company.
type IncomeStatement struct {
	Symbol           string              `json:"symbol"`
	AnnualReports    []IsAnnualReport    `json:"annualReports"`
	QuarterlyReports []IsQuarterlyReport `json:"quarterlyReports"`
}

// IsAnnualReport represents an annual report in the income statement.
type IsAnnualReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       AVInt  `json:"grossProfit"`
	TotalRevenue                      AVInt  `json:"totalRevenue"`
	CostOfRevenue                     AVInt  `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        AVInt  `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   AVInt  `json:"operatingIncome"`
	SellingGeneralAndAdministrative   AVInt  `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            AVInt  `json:"researchAndDevelopment"`
	OperatingExpenses                 AVInt  `json:"operatingExpenses"`
	InvestmentIncomeNet               AVInt  `json:"investmentIncomeNet"`
	NetInterestIncome                 AVInt  `json:"netInterestIncome"`
	InterestIncome                    AVInt  `json:"interestIncome"`
	InterestExpense                   AVInt  `json:"interestExpense"`
	NonInterestIncome                 AVInt  `json:"nonInterestIncome"`
	OtherNonOperatingIncome           AVInt  `json:"otherNonOperatingIncome"`
	Depreciation                      AVInt  `json:"depreciation"`
	DepreciationAndAmortization       AVInt  `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   AVInt  `json:"incomeBeforeTax"`
	IncomeTaxExpense                  AVInt  `json:"incomeTaxExpense"`
	InterestAndDebtExpense            AVInt  `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations AVInt  `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       AVInt  `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              AVInt  `json:"ebit"`
	Ebitda                            AVInt  `json:"ebitda"`
	NetIncome                         AVInt  `json:"netIncome"`
}

// IsQuarterlyReport represents a quarterly report in the income statement.
type IsQuarterlyReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       AVInt  `json:"grossProfit"`
	TotalRevenue                      AVInt  `json:"totalRevenue"`
	CostOfRevenue                     AVInt  `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        AVInt  `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   AVInt  `json:"operatingIncome"`
	SellingGeneralAndAdministrative   AVInt  `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            AVInt  `json:"researchAndDevelopment"`
	OperatingExpenses                 AVInt  `json:"operatingExpenses"`
	InvestmentIncomeNet               AVInt  `json:"investmentIncomeNet"`
	NetInterestIncome                 AVInt  `json:"netInterestIncome"`
	InterestIncome                    AVInt  `json:"interestIncome"`
	InterestExpense                   AVInt  `json:"interestExpense"`
	NonInterestIncome                 AVInt  `json:"nonInterestIncome"`
	OtherNonOperatingIncome           AVInt  `json:"otherNonOperatingIncome"`
	Depreciation                      AVInt  `json:"depreciation"`
	DepreciationAndAmortization       AVInt  `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   AVInt  `json:"incomeBeforeTax"`
	IncomeTaxExpense                  AVInt  `json:"incomeTaxExpense"`
	InterestAndDebtExpense            AVInt  `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations AVInt  `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       AVInt  `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              AVInt  `json:"ebit"`
	Ebitda                            AVInt  `json:"ebitda"`
	NetIncome                         AVInt  `json:"netIncome"`
}

func toIncomeStatement(buf []byte) (*IncomeStatement, error) {
	incomeStatement := &IncomeStatement{}
	if err := json.Unmarshal(buf, incomeStatement); err != nil {
		return nil, err
	}
	return incomeStatement, nil
}

// IncomeStatement fetches and returns the income statement data for the specified company symbol.
func (c *Client) IncomeStatement(symbol string) (*IncomeStatement, error) {
	const function = "INCOME_STATEMENT"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toIncomeStatement(body)
}
