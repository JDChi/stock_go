package china_money

import "testing"

func TestChinaMoneyData_GetRatesData(t *testing.T) {
	chinaMoneyData := NewChinaMoney()
	result, err := chinaMoneyData.GetRatesData()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", result)
}
