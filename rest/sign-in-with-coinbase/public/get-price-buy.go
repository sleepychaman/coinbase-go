package public

import (
	"net/http"
)

type RequestForGetPriceBuy struct {
	CurrencyPair string
}

type ResponseForGetPriceBuy struct {
	Data Price `json:"data"`
}

func (req *RequestForGetPriceBuy) Path() string {
	return "/v2/prices/" + req.CurrencyPair + "/buy"
}

func (req *RequestForGetPriceBuy) Method() string {
	return http.MethodGet
}

func (req *RequestForGetPriceBuy) Query() string {
	return ""
}

func (req *RequestForGetPriceBuy) Payload() []byte {
	return nil
}
