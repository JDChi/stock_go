package china_money

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"time"
)

func (c *ChinaMoney) GetLPRData() (*LPRData, error) {
	req := map[string]any{
		"t": time.Now().UnixMilli(),
	}
	reqJson, _ := json.Marshal(req)

	resp, err := http.Post(LPRUrl, "application/json", bytes.NewBuffer(reqJson))
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

	var ratesData *LPRData
	err = json.Unmarshal(body, &ratesData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return ratesData, nil
}

type LPRData struct {
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
		ShowDateEN string `json:"showDateEN"`
		ShowDateCN string `json:"showDateCN"`
	} `json:"data"`
	Records []struct {
		TermCode     string `json:"termCode"` // 期限 1Y 5Y
		Shibor       string `json:"shibor"`   // 贷款利率
		ShibIdUpDown string `json:"shibIdUpDown"`
	} `json:"records"`
}

// GetOneYearLPR 获取1年贷款利率
func (d *LPRData) GetOneYearLPR() float64 {
	for _, record := range d.Records {
		if record.TermCode == "1Y" {
			return cast.ToFloat64(record.Shibor) / 100
		}
	}
	return 0
}
