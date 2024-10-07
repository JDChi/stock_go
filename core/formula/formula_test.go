package formula

import "testing"

func TestCalPERatioOfTheStockMarket(t *testing.T) {
	result := CalPERatioOfTheStockMarket(0.04, 0.06)
	t.Logf("%+v", result)

}

func TestCalTheValuationOfTheStock(t *testing.T) {
	result := CalTheValuationOfTheStock(0.03, 0.5)
	t.Log(result)
}
