package crawler

import "testing"

func TestStockCrawler_CrawlerLPR(t *testing.T) {
	crawler := NewStockCrawler()
	result := crawler.CrawlerLPR()
	t.Log(result)
}
