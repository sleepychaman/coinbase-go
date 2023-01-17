package transactions

import (
	"encoding/json"
	"net/http"
)

type RequestForTransferMoney struct {
	AccountId   string            `json:"-"`
	To          string            `json:"to"`
	Type        TransferMoneyType `json:"type"`
	Amount      string            `json:"amount"`
	Currency    string            `json:"currency"`
	Description string            `json:"description,omitempty"`
}

type TransferMoneyType string

const (
	TransferType TransferMoneyType = "transfer"
)

type ResponseForTransferMoney struct {
	Data Transaction `json:"data"`
}

func (req *RequestForTransferMoney) Path() string {
	return "/v2/accounts/" + req.AccountId + "/transactions"
}

func (req *RequestForTransferMoney) Method() string {
	return http.MethodPost
}

func (req *RequestForTransferMoney) Query() string {
	return ""
}

func (req *RequestForTransferMoney) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
