package order

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForListOrders struct {
	ProductId            string                                `url:"product_id, omitempty"`
	OrderConfiguration   map[OrderTypeConfig]ConfigurationType `url:"order_configuration, omitempty"`
	OrderSide            OrderSide                             `url:"order_side, omitempty"`
	ClientOrderId        string                                `url:"client_order_id, omitempty"`
	Status               OrderStatus                           `url:"status, omitempty"`
	TimeInForce          string                                `url:"time_in_force, omitempty"`
	CreatedTime          time.Time                             `url:"created_time, omitempty"`
	CompletionPercentage string                                `url:"completion_percentage, omitempty"`
	FilledSize           string                                `url:"filled_size, omitempty"`
	AverageFilledPrice   string                                `url:"average_filled_price, omitempty"`
	Fee                  string                                `url:"fee, omitempty"`
	NumberOfFills        string                                `url:"number_of_fills, omitempty"`
	FilledValue          string                                `url:"filled_value, omitempty"`
	PendingCancel        bool                                  `url:"pending_cancel, omitempty"`
	SizeInQuote          bool                                  `url:"size_in_quote, omitempty"`
	TotalFees            string                                `url:"total_fees, omitempty"`
}

type ResponseForListOrders struct {
	Orders []Order `json:"orders,omitempty"`
}

func (req *RequestForListOrders) Path() string {
	return "/api/v3/brokerage/orders/historical/batch/"
}

func (req *RequestForListOrders) Method() string {
	return http.MethodGet
}

func (req *RequestForListOrders) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForListOrders) Payload() []byte {
	return nil
}
