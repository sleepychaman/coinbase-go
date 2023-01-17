package products

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/sleepychaman/coinbase-go/advanded-trade-api/products"
)

type RequestForListProducts struct {
	StartDate          time.Time            `url:"start_date,omitempty"`
	EndDate            time.Time            `url:"end_date,omitempty"`
	UserNaticeCurrency string               `url:"user_native_currency,omitempty"`
	ProductType        products.ProductType `url:"product_type,omitempty"`
}

type ResponseForListProducts struct {
	TotalVolume             float64                 `json:"total_volume"`
	TotalFees               float64                 `json:"total_fees"`
	FeeTier                 FeeTierType             `json:"fee_tier"`
	MarginRate              MarginRateType          `json:"margin_rate"`
	GoodsAndServicesTax     GoodsAndServicesTaxType `json:"goods_and_services_tax"`
	AdvancedTradeOnlyVolume float64                 `json:"advanced_trade_only_volume"`
	AdvandedTradeOnlyFees   float64                 `json:"advanced_trade_only_fees"`
	CoinbaseProVolume       float64                 `json:"coinbase_pro_volume"`
	CoinbaseProFees         float64                 `json:"coinbase_pro_fees"`
}

type FeeTierType struct {
	PricingTier  string `json:"pricing_tier"`
	UsdFrom      string `json:"usd_from"`
	UsdTo        string `json:"usd_to"`
	TakerFeeRate string `json:"taker_fee_rate"`
	MakerFeeRate string `json:"maker_fee_rate"`
}

type MarginRateType struct {
	Value string `json:"value"`
}

type GoodsAndServicesTaxType struct {
	Rate string `json:"rate"`
	Type string `json:"type"`
}

func (req *RequestForListProducts) Path() string {
	return "/api/v3/brokerage/transaction_summary/"
}

func (req *RequestForListProducts) Method() string {
	return http.MethodGet
}

func (req *RequestForListProducts) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForListProducts) Payload() []byte {
	return nil
}
