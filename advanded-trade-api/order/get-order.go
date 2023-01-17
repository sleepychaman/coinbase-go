package order

import (
	"net/http"
	"time"
)

type RequestForGetOrder struct {
	OrderId string `json:"order_id"`
}

type ResponseForGetOrder struct {
	Order Order `json:"order"`
}

type Order struct {
	OrderId              string                                `json:"order_id"`
	ProductID            string                                `json:"product_id"`
	UserId               string                                `json:"user_id"`
	OrderConfiguration   map[OrderTypeConfig]ConfigurationType `json:"order_configuration"`
	Side                 OrderSide                             `json:"side"`
	ClientOrderId        string                                `json:"client_order_id"`
	Status               OrderStatus                           `json:"status"`
	TimeInForce          string                                `json:"time_in_force"`
	CreatedTime          time.Time                             `json:"created_time"`
	CompletionPercentage string                                `json:"completion_percentage"`
	FilledSize           string                                `json:"filled_size"`
	AverageFilledPrice   string                                `json:"average_filled_price"`
	Fee                  string                                `json:"fee"`
	NumberOfFills        string                                `json:"number_of_fills"`
	FilledValue          string                                `json:"filled_value"`
	PendingCancel        bool                                  `json:"pending_cancel"`
	SizeInQuote          bool                                  `json:"size_in_quote"`
	TotalFees            string                                `json:"total_fees"`
	SizeInclusiveOfFees  bool                                  `json:"size_inclusive_of_fees"`
	TotalValueAfterFees  string                                `json:"total_value_after_fees"`
	TriggerStatus        string                                `json:"trigger_status"`
	OrderType            string                                `json:"order_type"`
	RejectReason         string                                `json:"reject_reason"`
	Settled              bool                                  `json:"settled"`
	ProductType          string                                `json:"product_type"`
	RejectMessage        string                                `json:"reject_message"`
	CancelMessage        string                                `json:"cancel_message"`
}

type OrderStatus string

const (
	OPEN_STATUS          OrderStatus = "OPEN"
	FILLED_STATUS        OrderStatus = "FILLED"
	CANCELLED_STATUS     OrderStatus = "CANCELLED"
	EXPIRED_STATUS       OrderStatus = "EXPIRED"
	FAILED_STATUS        OrderStatus = "FAILED"
	UNKNOWN_ORDER_STATUS OrderStatus = "UNKNOWN_ORDER_STATUS"
)

type OrderSide string

const (
	UNKNOWN_ORDER_SIDE OrderSide = "UNKNOWN_ORDER_SIDE"
	BUY_ORDER_SIDE     OrderSide = "BUY"
	SELL_ORDER_SIDE    OrderSide = "SELL"
)

func (req *RequestForGetOrder) Path() string {
	return "/api/v3/brokerage/orders/historical/" + req.OrderId
}

func (req *RequestForGetOrder) Method() string {
	return http.MethodGet
}

func (req *RequestForGetOrder) Query() string {
	return ""
}

func (req *RequestForGetOrder) Payload() []byte {
	return nil
}
