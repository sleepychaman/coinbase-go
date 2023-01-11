package products

import (
	"net/http"
)

type RequestForGetProduct struct {
	ProductID string
}

type ResponseForGetProduct struct {
	ProductID                 string `json:"product_id"`
	Price                     string `json:"price"`
	PricePercentageChange24H  string `json:"price_percentage_change_24h"`
	Volume24H                 string `json:"volume_24h"`
	VolumePercentageChange24H string `json:"volume_percentage_change_24h"`
	BaseIncrement             string `json:"base_increment"`
	QuoteIncrement            string `json:"quote_increment"`
	QuoteMinSize              string `json:"quote_min_size"`
	QuoteMaxSize              string `json:"quote_max_size"`
	BaseMinSize               string `json:"base_min_size"`
	BaseMaxSize               string `json:"base_max_size"`
	BaseName                  string `json:"base_name"`
	QuoteName                 string `json:"quote_name"`
	Watched                   bool   `json:"watched"`
	Isdisabled                bool   `json:"is_disabled"`
	New                       bool   `json:"new"`
	Status                    string `json:"status"`
	CancelOnly                bool   `json:"cancel_only"`
	LimitOnly                 bool   `json:"limit_only"`
	PostOnly                  bool   `json:"post_only"`
	TradingDisabled           bool   `json:"trading_disabled"`
	AuctionMode               bool   `json:"auction_mode"`
	ProductType               string `json:"product_type"`
	QuoteCurrencyID           string `json:"quote_currency_id"`
	BaseCurrencyID            string `json:"base_currency_id"`
	MidMarketPrice            string `json:"mid_market_price"`
}

type Product struct{}

func (req *RequestForGetProduct) Path() string {
	return "/api/v3/brokerage/products/" + req.ProductID
}

func (req *RequestForGetProduct) Method() string {
	return http.MethodGet
}

func (req *RequestForGetProduct) Query() string {
	return ""
}

func (req *RequestForGetProduct) Payload() []byte {
	return nil
}
