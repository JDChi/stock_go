package stock

import (
	"context"
	"my_stock/core/formula"
)

// CalTheValuationOfTheStock 计算股票的估值
func CalTheValuationOfTheStock(ctx context.Context, secuCode string) {
	// TODO 获取股票最近一次的分红价格
	// TODO 获取银行的一年定期存款利率
	// TODO 计算股票的估值
	formula.CalTheValuationOfTheStock(0.05, 0.06)
}
