package transactions

import (
	"net/http"
	"time"

	"github.com/sleepychaman/coinbase-go/advanded-trade-api/account"
)

type RequestForGetTransaction struct {
	AccountId     string
	TransactionId string
}

type ResponseForGetTransaction struct {
	Data Transaction `json:"data"`
}

type Transaction struct {
	Id           string            `json:"id"`
	Type         string            `json:"type"`
	Status       account.Balance   `json:"status"`
	Amount       account.Balance   `json:"amount"`
	NativeAmount string            `json:"native_amount"`
	Description  string            `json:"description"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Resource     string            `json:"resource"`
	ResourcePath string            `json:"resource_path"`
	To           map[string]string `json:"to"`
	Details      map[string]string `json:"details"`
}

func (req *RequestForGetTransaction) Path() string {
	return "/v2/accounts/" + req.AccountId + "/transactions/" + req.TransactionId
}

func (req *RequestForGetTransaction) Method() string {
	return http.MethodGet
}

func (req *RequestForGetTransaction) Query() string {
	return ""
}

func (req *RequestForGetTransaction) Payload() []byte {
	return nil
}
