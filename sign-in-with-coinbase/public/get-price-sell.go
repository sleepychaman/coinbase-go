package public

import (
	"net/http"
)

type RequestForGetPriceSell struct {
	CurrencyPair string
}

type ResponseForGetPriceSell struct {
	Data Price `json:"data"`
}

func (req *RequestForGetPriceSell) Path() string {
	return "/v2/prices/" + req.CurrencyPair + "/sell"
}

func (req *RequestForGetPriceSell) Method() string {
	return http.MethodGet
}

func (req *RequestForGetPriceSell) Query() string {
	return ""
}

func (req *RequestForGetPriceSell) Payload() []byte {
	return nil
}
