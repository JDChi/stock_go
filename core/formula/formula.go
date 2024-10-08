package formula

import "encoding/json"

// 记录一些计算公式

// CalPERatioOfTheStockMarket 计算股市整体市盈率
//
// 股市整体市盈率 = 1 / 利率，其中利率看个人自己的取值，按照《投资大白话》一书，利率是取（银行一年期存款利率 + 银行一年期贷款利率） / 2
// 然后分别估算最低市盈率和最高市盈率，可以跟实际的值进行比较，认为低于最低市盈率，是低估，高于最高市盈率，是高估
func CalPERatioOfTheStockMarket(depositInterestRate, loanInterestRate float64) StockMarketPERatioResult {

	result := StockMarketPERatioResult{}
	result.ReasonablePERatio = 1 / ((depositInterestRate + loanInterestRate) / 2)
	if depositInterestRate < loanInterestRate {
		result.MaxPERatio = 1 / depositInterestRate
		result.MinPERatio = 1 / loanInterestRate
	} else {
		result.MaxPERatio = 1 / loanInterestRate
		result.MinPERatio = 1 / depositInterestRate
	}

	return result
}

// StockMarketPERatioResult 股市整体市盈率计算结果
type StockMarketPERatioResult struct {
	// 最低市盈率
	MinPERatio float64 `json:"min_pe_ratio"`
	// 最高市盈率
	MaxPERatio float64 `json:"max_pe_ratio"`
	// 合理市盈率
	ReasonablePERatio float64 `json:"reasonable_pe_ratio"`
}

func (s *StockMarketPERatioResult) String() string {
	m, _ := json.Marshal(s)
	return string(m)
}

// CalTheValuationOfTheStock 计算单只股票的估值
//
// 这里使用股息率和银行存款利率来给个股估值。也就是通过个股的分红价格，以及银行存款利率来反推个股的股价，来自《投资大白话》一书
func CalTheValuationOfTheStock(dividendPrice float64, depositInterestRate float64) float64 {
	return dividendPrice / depositInterestRate
}
