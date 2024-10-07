package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cast"

	"strings"
)

// CrawlFixedDepositedInterestRate 爬取存款利率
func (c *StockCrawler) CrawlFixedDepositedInterestRate() FixedDepositedInterestRateResult {
	// 是否位于整存整取存款利率之后
	isInnerFixedDeposit := false
	result := FixedDepositedInterestRateResult{}

	c.collector.OnHTML("tr", func(element *colly.HTMLElement) {
		// 需要确定数据的位置
		if strings.Contains(element.Text, "整存整取定期存款") {
			isInnerFixedDeposit = true
			return
		}

		if isInnerFixedDeposit {
			// 获取当前行的所有 <td> 标签
			cols := element.ChildTexts("td")

			if len(cols) >= 3 {
				isThreeMonths := strings.Contains(element.Text, "三个月")
				isHalfYear := strings.Contains(element.Text, "半年")
				isOneYear := strings.Contains(element.Text, "一年")
				isTwoYears := strings.Contains(element.Text, "二年")
				isThreeYears := strings.Contains(element.Text, "三年")
				isFiveYears := strings.Contains(element.Text, "五年")

				if isThreeMonths {
					result.ThreeMonths = cast.ToFloat64(cols[1])
				} else if isHalfYear {
					result.HalfYear = cast.ToFloat64(cols[1])
				} else if isOneYear {
					result.OneYear = cast.ToFloat64(cols[1])
				} else if isTwoYears {
					result.TwoYears = cast.ToFloat64(cols[1])
				} else if isThreeYears {
					result.ThreeYears = cast.ToFloat64(cols[1])
				} else if isFiveYears {
					result.FiveYears = cast.ToFloat64(cols[1])
				}
			}
		}

	})

	_ = c.collector.Visit(InterestRateUrl)

	return result
}

// FixedDepositedInterestRateResult 存款利率结果
type FixedDepositedInterestRateResult struct {
	ThreeMonths float64 `json:"three-months"`
	HalfYear    float64 `json:"half-year"`
	OneYear     float64 `json:"one-year"`
	TwoYears    float64 `json:"two-years"`
	ThreeYears  float64 `json:"three-years"`
	FiveYears   float64 `json:"five-years"`
}
