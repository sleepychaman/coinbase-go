package order

import (
	"net/http"
	"time"

	"github.com/sleepychaman/coinbase-go/rest/generic"
)

type RequestForListFills struct {
	generic.Pagination
	OrderID                string    `url:"order_id"`
	ProductID              string    `url:"product_id"`
	StartSequenceTimestamp time.Time `url:"start_sequence_timestamp"`
	EndSequenceTimestamp   time.Time `url:"end_sequence_timestamp"`
}

type ResponseForListFills struct {
	generic.Pagination
	Fills []Fill `json:"fills"`
}

type Fill struct {
	EntryID            string                 `json:"entry_id"`
	TradeID            string                 `json:"trade_id"`
	OrderID            string                 `json:"order_id"`
	TradeTime          time.Time              `json:"trade_time"`
	TradeType          FillType               `json:"trade_type"`
	Price              string                 `json:"price"`
	Size               string                 `json:"size"`
	Commission         string                 `json:"commission"`
	ProductID          string                 `json:"product_id"`
	SequenceTimestamp  time.Time              `json:"sequence_timestamp"`
	LiquidityIndicator LiquidityIndicatorType `json:"liquidity_indicator"`
	SizeInQuote        bool                   `json:"size_in_quote"`
	UserID             string                 `json:"user_id"`
	Side               OrderSide              `json:"side"`
}

type LiquidityIndicatorType string

const (
	UNKNOWN_LIQUIDITY_INDICATOR_TYPE LiquidityIndicatorType = "UNKNOWN_LIQUIDITY_INDICATOR"
	MAKER_TYPE                       LiquidityIndicatorType = "MAKER"
	TAKER_TYPE                       LiquidityIndicatorType = "TAKER"
)

type FillType string

const (
	FILL_TYPE       FillType = "FILL"
	REVERSAL_TYPE   FillType = "REVERSAL"
	CORRECTION_TYPE FillType = "CORRECTION"
	SYNTHETIC_TYPE  FillType = "SYNTHETIC"
)

func (req *RequestForListFills) Path() string {
	return "/api/v3/brokerage/orders/historical/fills/"
}

func (req *RequestForListFills) Method() string {
	return http.MethodGet
}

func (req *RequestForListFills) Query() string {
	return ""
}

func (req *RequestForListFills) Payload() []byte {
	return nil
}
