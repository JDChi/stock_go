package east_money

type EastMoney struct {
}

func NewEastMoney() *EastMoney {
	return &EastMoney{}
}

type MarketValues struct {
	Version string                 `json:"version"`
	Result  StockMarketValueResult `json:"result"`
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Code    int                    `json:"code"`
}

func (s *MarketValues) GetLatestData() StockMarketValueData {
	if len(s.Result.Data) == 0 {
		return StockMarketValueData{}
	}
	return s.Result.Data[0]
}

type StockMarketValueResult struct {
	Pages int                    `json:"pages"`
	Data  []StockMarketValueData `json:"data"`
	Count int                    `json:"count"`
}

type StockMarketValueData struct {
	TRADEMARKETCODE   string  `json:"TRADE_MARKET_CODE"`
	TRADEDATE         string  `json:"TRADE_DATE"`
	CLOSEPRICE        float64 `json:"CLOSE_PRICE"`
	CHANGERATE        float64 `json:"CHANGE_RATE"`
	SECURITYINNERCODE string  `json:"SECURITY_INNER_CODE"`
	LISTINGORGNUM     int     `json:"LISTING_ORG_NUM"`
	TOTALSHARES       float64 `json:"TOTAL_SHARES"`
	FREESHARES        float64 `json:"FREE_SHARES"`
	TOTALMARKETCAP    float64 `json:"TOTAL_MARKET_CAP"`
	FREEMARKETCAP     float64 `json:"FREE_MARKET_CAP"`
	PETTMAVG          float64 `json:"PE_TTM_AVG"` // 平均市盈率
}
