package order

import (
	"encoding/json"
	"net/http"
	"time"
)

type RequestForCreateOrder struct {
	ClientOrderID      string                                `json:"client_order_id"`
	ProductID          string                                `json:"product_id"`
	Side               OrderSide                             `json:"side"`
	OrderConfiguration map[OrderTypeConfig]ConfigurationType `json:"order_configuration,omitempty"`
}

type ResponseForCreateOrder struct {
	Success            bool                                  `json:"success"`
	FailureReason      string                                `json:"failure_reason,omitempty"`
	OrderId            string                                `json:"order_id"`
	SuccessResponse    SuccessResponseType                   `json:"success_response,omitempty"`
	ErrorResponse      ErrorResponseType                     `json:"error_response,omitempty"`
	OrderConfiguration map[OrderTypeConfig]ConfigurationType `json:"order_configuration,omitempty"`
}

type SuccessResponseType struct {
	OrderId       string    `json:"order_id"`
	ProductID     string    `json:"product_id"`
	Side          OrderSide `json:"side"`
	ClientOrderID string    `json:"client_order_id"`
}
type ErrorResponseType struct {
	Error                 string `json:"error"`
	Message               string `json:"message"`
	ErrorDetails          string `json:"error_details"`
	PreviewFailureReason  string `json:"preview_failure_reason"`
	NewOrderFailureReason string `json:"new_order_failure_reason"`
}

type OrderTypeConfig string

const (
	MarketMarketIoc       OrderTypeConfig = "market_market_ioc"
	LimitLimitGtc         OrderTypeConfig = "limit_limit_gtc"
	LimitLimitGtd         OrderTypeConfig = "limit_limit_gtd"
	StopLimitStopLimitGtc OrderTypeConfig = "stop_limit_stop_limit_gtc"
	StopLimitStopLimitGtd OrderTypeConfig = "stop_limit_stop_limit_gtd"
)

type ConfigurationType struct {
	QuoteSize     string     `json:"quote_size,omitempty"`
	BaseSize      string     `json:"base_size,omitempty"`
	LimitPrice    string     `json:"limit_price,omitempty"`
	PostOnly      bool       `json:"post_only,omitempty"`
	EndTime       *time.Time `json:"end_time,omitempty"`
	StopDirection string     `json:"stop_direction,omitempty"`
	StopPrice     string     `json:"stop_price,omitempty"`
}

func (req *RequestForCreateOrder) Path() string {
	return "/api/v3/brokerage/orders/"
}

func (req *RequestForCreateOrder) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateOrder) Query() string {
	return ""
}

func (req *RequestForCreateOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
