package transactions

import (
	"encoding/json"
	"net/http"
)

type RequestForSendMoney struct {
	AccountId                   string        `json:"-"`
	To                          string        `json:"to"`
	Type                        SendMoneyType `json:"type"`
	Amount                      string        `json:"amount"`
	Currency                    string        `json:"currency"`
	Description                 string        `json:"description,omitempty"`
	Skipnotifications           bool          `json:"skip_notifications,omitempty"`
	Idem                        string        `json:"idem,omitempty"`
	ToFinancialInstitution      string        `json:"to_financial_institution,omitempty"`
	FinancialInstitutionWebsite string        `json:"financial_institution_website,omitempty"`
}

type SendMoneyType string

const (
	SendType SendMoneyType = "send"
)

type ResponseForSendMoney struct {
	Data Transaction `json:"data"`
}

func (req *RequestForSendMoney) Path() string {
	return "/v2/accounts/" + req.AccountId + "/transactions"
}

func (req *RequestForSendMoney) Method() string {
	return http.MethodPost
}

func (req *RequestForSendMoney) Query() string {
	return ""
}

func (req *RequestForSendMoney) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
