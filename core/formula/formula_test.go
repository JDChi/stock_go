package formula

import "testing"

func TestCalPERatioOfTheStockMarket(t *testing.T) {
	result := CalPERatioOfTheStockMarket(0.0135, 0.0335)
	t.Logf("min: %f", result.MinPERatio)
	t.Logf("max: %f", result.MaxPERatio)
	t.Logf("reasonable: %f", result.ReasonablePERatio)
	t.Logf("current SH: %f", 14.24)
	t.Logf("current SZ: %f", 23.68)
}

func TestCalTheValuationOfTheStock(t *testing.T) {
	result := CalTheValuationOfTheStock(1.3, 0.0135)
	t.Logf("value: %f", result)
}
