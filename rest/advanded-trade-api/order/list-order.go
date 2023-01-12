package order

import (
	"net/http"
	"time"
)

type RequestForListOrder struct {
	ProductId            string                                `json:"product_id"`
	OrderConfiguration   map[OrderTypeConfig]ConfigurationType `json:"order_configuration"`
	OrderSide            OrderSide                             `json:"order_side"`
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
}

type ResponseForListOrder struct {
	Orders []Order `json:"orders"`
}

func (req *RequestForListOrder) Path() string {
	return "/api/v3/brokerage/orders/historical/batch/"
}

func (req *RequestForListOrder) Method() string {
	return http.MethodGet
}

func (req *RequestForListOrder) Query() string {
	return ""
}

func (req *RequestForListOrder) Payload() []byte {
	return nil
}
