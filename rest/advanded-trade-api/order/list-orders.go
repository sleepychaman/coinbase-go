package order

import (
	"net/http"
	"time"
)

type RequestForListOrders struct {
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

type ResponseForListOrders struct {
	Orders []Order `json:"orders"`
}

func (req *RequestForListOrders) Path() string {
	return "/api/v3/brokerage/orders/historical/batch/"
}

func (req *RequestForListOrders) Method() string {
	return http.MethodGet
}

func (req *RequestForListOrders) Query() string {
	return ""
}

func (req *RequestForListOrders) Payload() []byte {
	return nil
}
