package east_money

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetSHValues 获取沪市估值数据
func (e *EastMoney) GetSHValues() (*MarketValues, error) {
	params := url.Values{}
	params.Set("sortColumns", "TRADE_DATE")
	params.Set("sortTypes", "-1")
	params.Set("pageSize", "10")
	params.Set("pageNumber", "1")
	params.Set("reportName", "RPT_VALUEMARKET")
	params.Set("columns", "ALL")
	params.Set("source", "WEB")
	params.Set("client", "WEB")
	params.Set("filter", "(TRADE_MARKET_CODE=000001)")

	// 将参数附加到 URL
	url := fmt.Sprintf("%s?%s", MarketValuesUrl, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
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

	var marketValues *MarketValues
	err = json.Unmarshal(body, &marketValues)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return marketValues, nil

}
