package crawler

import (
	"crypto/tls"
	"github.com/gocolly/colly/v2"
	"net/http"
)

type StockCrawler struct {
	collector *colly.Collector
}

func NewStockCrawler() *StockCrawler {
	collector := colly.NewCollector()
	collector.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	return &StockCrawler{
		collector: collector,
	}
}
