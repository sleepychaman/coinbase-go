package users

import (
	"net/http"
)

type RequestForShowCurrentUser struct{}

type ResponseForShowCurrentUser struct {
	Data User `json:"data"`
}

func (req *RequestForShowCurrentUser) Path() string {
	return "/v2/user"
}

func (req *RequestForShowCurrentUser) Method() string {
	return http.MethodGet
}

func (req *RequestForShowCurrentUser) Query() string {
	return ""
}

func (req *RequestForShowCurrentUser) Payload() []byte {
	return nil
}
