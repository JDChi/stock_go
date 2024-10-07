package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

// CrawlerLPR 爬取 LPR 贷款市场报价利率
func (c *StockCrawler) CrawlerLPR() LPRResult {

	isInnerLPR := false
	result := LPRResult{}

	c.collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	c.collector.OnHTML("tr", func(element *colly.HTMLElement) {
		cols := element.ChildTexts("td")

		if strings.Contains(element.Text, "LPR报价") {
			isInnerLPR = true
			return
		}

		if isInnerLPR {
			if len(cols) >= 3 {
				fmt.Println(cols)
			}

		}

	})
	err := c.collector.Visit(LPRUrl)
	if err != nil {
		fmt.Println(err)
	}

	return result

}

type LPRResult struct {
	OneYear   float64 `json:"one-year"`
	FiveYears float64 `json:"five-years"`
}
