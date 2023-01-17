package products

import (
	"net/http"
)

type RequestForGetProductCandles struct {
	ProductID   string          `url:"-"`
	Start       string          `url:"start"`
	End         string          `url:"end"`
	Granularity GranularityType `url:"granularity"`
}

type ResponseForGetProductCandles struct {
	Candles []ProductCandles `json:"candles"`
}

type ProductCandles struct {
	Start  string `json:"start"`
	Low    string `json:"low"`
	High   string `json:"high"`
	Open   string `json:"open"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
}

type GranularityType string

const (
	UNKNOWN_GRANULARITY_TYPE GranularityType = "UNKNOWN_GRANULARITY"
	ONE_MINUTE_TYPE          GranularityType = "ONE_MINUTE"
	FIVE_MINUTE_TYPE         GranularityType = "FIVE_MINUTE"
	FIFTEEN_MINUTE_TYPE      GranularityType = "FIFTEEN_MINUTE"
	THIRTY_MINUTE_TYPE       GranularityType = "THIRTY_MINUTE"
	ONE_HOUR_TYPE            GranularityType = "ONE_HOUR"
	TWO_HOUR_TYPE            GranularityType = "TWO_HOUR"
	SIX_HOUR_TYPE            GranularityType = "SIX_HOUR"
	ONE_DAY_TYPE             GranularityType = "ONE_DAY"
)

func (req *RequestForGetProductCandles) Path() string {
	return "/api/v3/brokerage/products/" + req.ProductID + "/candles/"
}

func (req *RequestForGetProductCandles) Method() string {
	return http.MethodGet
}

func (req *RequestForGetProductCandles) Query() string {
	return ""
}

func (req *RequestForGetProductCandles) Payload() []byte {
	return nil
}
