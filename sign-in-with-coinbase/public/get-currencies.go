package public

import "net/http"

type RequestForGetCurrencies struct{}

type ResponseForGetCurrencies struct {
	Data []CurrencyInfo `json:"data"`
}

type CurrencyInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
}

func (req *RequestForGetCurrencies) Path() string {
	return "api/v2/currencies"
}

func (req *RequestForGetCurrencies) Method() string {
	return http.MethodGet
}

func (req *RequestForGetCurrencies) Query() string {
	return ""
}

func (req *RequestForGetCurrencies) Payload() []byte {
	return nil
}
