package public

import (
	"net/http"
)

type RequestForGetPriceSpot struct {
	CurrencyPair string
}

type ResponseForGetPriceSpot struct {
	Data Price `json:"data"`
}

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

func (req *RequestForGetPriceSpot) Path() string {
	return "/v2/prices/" + req.CurrencyPair + "/spot"
}

func (req *RequestForGetPriceSpot) Method() string {
	return http.MethodGet
}

func (req *RequestForGetPriceSpot) Query() string {
	return ""
}

func (req *RequestForGetPriceSpot) Payload() []byte {
	return nil
}
