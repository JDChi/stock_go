package east_money

import "testing"

func TestEastMoney_GetSHValues(t *testing.T) {
	e := NewEastMoney()
	result, err := e.GetSHValues()
	if err != nil {
		t.Errorf("GetSZValues() error = %v", err)
		return
	}

	t.Logf("result = %+v", result)
	t.Logf("latestData: %+v", result.GetLatestData())
}
