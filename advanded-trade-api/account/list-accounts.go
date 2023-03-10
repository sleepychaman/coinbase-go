package account

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/sleepychaman/coinbase-go/generic"
)

type RequestForListAccounts struct {
	generic.AdvancedTradePaginationRequest
}

type ResponseForListAccounts struct {
	generic.AdvancedTradePaginationResponse
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Uuid             string      `json:"uuid"`
	Name             string      `json:"name"`
	Currency         string      `json:"currency"`
	AvailableBalance Balance     `json:"available_balance"`
	Default          bool        `json:"default"`
	Active           bool        `json:"active"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	DeletedAt        time.Time   `json:"deleted_at"`
	Type             AccountType `json:"type"`
	Ready            bool        `json:"ready"`
	Hold             Balance     `json:"hold"`
}
type AccountType string

const (
	ACCOUNT_TYPE_CRYPTO AccountType = "ACCOUNT_TYPE_CRYPTO"
	ACCOUNT_TYPE_FIAT   AccountType = "ACCOUNT_TYPE_FIAT"
)

type Balance struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

func (req *RequestForListAccounts) Path() string {
	return "/api/v3/brokerage/accounts"
}

func (req *RequestForListAccounts) Method() string {
	return http.MethodGet
}

func (req *RequestForListAccounts) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForListAccounts) Payload() []byte {
	return nil
}
