package china_money

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ChinaMoneyData struct {
}

func NewChinaMoneyData() *ChinaMoneyData {
	return &ChinaMoneyData{}
}

// GetRatesData 获取利率数据
func (c *ChinaMoneyData) GetRatesData() (*RatesData, error) {
	req := ratesDataRequest{
		T: time.Now().UnixMilli(),
	}
	reqJson, _ := json.Marshal(req)

	resp, err := http.Post(RateDataUrl, "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	var respData *RatesData
	err = json.Unmarshal(body, &respData)

	return respData, nil
}

type ratesDataRequest struct {
	T int64 `json:"t"`
}

type RatesData struct {
	Head struct {
		Version    string `json:"version"`
		Provider   string `json:"provider"`
		ReqCode    string `json:"req_code"`
		RepCode    string `json:"rep_code"`
		RepMessage string `json:"rep_message"`
		Ts         int64  `json:"ts"`
		Producer   string `json:"producer"`
	} `json:"head"`
	Data struct {
		Bond3M       string `json:"bond3M"`  // 3个月期国债发行利率
		Bond10Y      string `json:"bond10Y"` // 10 年期国债收益率
		LastDate3MEn string `json:"lastDate3MEn"`
		ShowDateEN   string `json:"showDateEN"`
		Bond1Y       string `json:"bond1Y"`     // 1 年期国债收益率
		DepoRate1Y   string `json:"depoRate1Y"` // 1 年期存款利率
		LastDate3M   string `json:"lastDate3M"` // 3个月期国债发行利率发行日期
		ShowDateCN   string `json:"showDateCN"`
		LendRate1Y   string `json:"lendRate1Y"` // 1 年期贷款利率
	} `json:"data"`
}
