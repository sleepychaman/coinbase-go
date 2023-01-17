package products

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/sleepychaman/coinbase-go/advanded-trade-api/order"
)

type RequestForGetMarketTrades struct {
	ProductID string `url:"-"`
	Limit     int32  `url:"limit"`
}

type ResponseForGetMarketTrades struct {
	Trades []ProductTicker `json:"trades"`
}

type ProductTicker struct {
	TradeID   string          `json:"trade_id"`
	ProductID string          `json:"product_id"`
	Price     string          `json:"price"`
	Size      string          `json:"size"`
	Time      time.Time       `json:"time"`
	Volume    string          `json:"volume"`
	Side      order.OrderSide `json:"side"`
	Bid       string          `json:"bid"`
	Ask       string          `json:"ask"`
}

func (req *RequestForGetMarketTrades) Path() string {
	return "/api/v3/brokerage/products/" + req.ProductID + "/ticker/"
}

func (req *RequestForGetMarketTrades) Method() string {
	return http.MethodGet
}

func (req *RequestForGetMarketTrades) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForGetMarketTrades) Payload() []byte {
	return nil
}
