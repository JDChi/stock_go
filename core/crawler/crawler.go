package crawler

import "github.com/gocolly/colly/v2"

type StockCrawler struct {
	collector *colly.Collector
}

func NewStockCrawler() *StockCrawler {
	colly.NewCollector()
	return &StockCrawler{
		collector: colly.NewCollector(),
	}
}
