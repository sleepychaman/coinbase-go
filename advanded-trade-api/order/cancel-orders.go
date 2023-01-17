package order

import (
	"encoding/json"
	"net/http"
)

type RequestForCancelOrders struct {
	OrderIds []string `json:"order_ids"`
}

type ResponseForCancelOrders struct {
	Results []ResultType `json:"results"`
}

type ResultType struct {
	Success       bool   `json:"success"`
	FailureReason string `json:"failure_reason,omitempty"`
	OrderId       string `json:"order_id"`
}

func (req *RequestForCancelOrders) Path() string {
	return "/api/v3/brokerage/orders/batch_cancel/"
}

func (req *RequestForCancelOrders) Method() string {
	return http.MethodPost
}

func (req *RequestForCancelOrders) Query() string {
	return ""
}

func (req *RequestForCancelOrders) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
