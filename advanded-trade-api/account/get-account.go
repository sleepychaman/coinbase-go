package account

import (
	"net/http"
)

type RequestForGetAccount struct {
	Uuid string
}

type ResponseForGetAccount struct {
	Account Account `json:"account"`
}

func (req *RequestForGetAccount) Path() string {
	return "/api/v3/brokerage/accounts/" + req.Uuid
}

func (req *RequestForGetAccount) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAccount) Query() string {
	return ""
}

func (req *RequestForGetAccount) Payload() []byte {
	return nil
}
