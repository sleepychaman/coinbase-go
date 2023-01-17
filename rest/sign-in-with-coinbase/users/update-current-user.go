package users

import (
	"encoding/json"
	"net/http"
)

type RequestForUpdateCurrentUser struct {
	Name           string `json:"name,omitempty"`
	TimeZone       string `json:"time_zone,omitempty"`
	NativeCurrency string `json:"native_currency,omitempty"`
}

type ResponseForUpdateCurrentUser struct {
	Data User `json:"data"`
}

func (req *RequestForUpdateCurrentUser) Path() string {
	return "/v2/user"
}

func (req *RequestForUpdateCurrentUser) Method() string {
	return http.MethodPut
}

func (req *RequestForUpdateCurrentUser) Query() string {
	return ""
}

func (req *RequestForUpdateCurrentUser) Payload() []byte {
	d, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return d
}
