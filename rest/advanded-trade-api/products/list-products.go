package products

import (
	"net/http"
)

type RequestForListProducts struct {
	Limit       int32       `url:"limit,omitempty"`
	Offset      int32       `url:"offset,omitempty"`
	ProductType ProductType `url:"product_type,omitempty"`
}

type ResponseForListProducts struct {
	Products    []Product `json:"products"`
	NumProducts int32     `json:"num_products"`
}

type ProductType string

const (
	SPOT_TYPE ProductType = "SPOT"
)

func (req *RequestForListProducts) Path() string {
	return "/api/v3/brokerage/products/"
}

func (req *RequestForListProducts) Method() string {
	return http.MethodGet
}

func (req *RequestForListProducts) Query() string {
	return ""
}

func (req *RequestForListProducts) Payload() []byte {
	return nil
}
