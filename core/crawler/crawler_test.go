package crawler

import "testing"

func TestStockCrawler_CrawlInterestRate(t *testing.T) {
	crawler := NewStockCrawler()
	result := crawler.CrawlFixedDepositedInterestRate()
	t.Logf("%+v", result)
}
