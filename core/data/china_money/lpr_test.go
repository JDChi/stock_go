package china_money

import "testing"

func TestChinaMoneyData_GetLPRData(t *testing.T) {
	chinaMoneyData := NewChinaMoney()
	lprData, err := chinaMoneyData.GetLPRData()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", lprData)
	t.Logf("one year lpr: %f", lprData.GetOneYearLPR())
}
