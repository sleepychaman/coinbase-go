package transactions

import (
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/sleepychaman/coinbase-go/generic"
)

type RequestForListTransactions struct {
	generic.SignInWithCoinbasePaginationRequest
	AccountId string `url:"-"`
}

type ResponseForListTransactions struct {
	generic.SignInWithCoinbasePaginationResponse
	Data []Transaction `json:"data"`
}

func (req *RequestForListTransactions) Path() string {
	return "/v2/accounts/" + req.AccountId + "/transactions"
}

func (req *RequestForListTransactions) Method() string {
	return http.MethodGet
}

func (req *RequestForListTransactions) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForListTransactions) Payload() []byte {
	return nil
}
