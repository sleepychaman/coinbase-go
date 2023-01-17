package users

import (
	"net/http"
)

type RequestForShowAuthorizationInfo struct{}

type ResponseForShowAuthorizationInfo struct {
	Data AuthorizationInfo `json:"data"`
}

type AuthorizationInfo struct {
	Method    string            `json:"method"`
	Scopes    []string          `json:"scopes"`
	OauthMeta map[string]string `json:"oauth_meta"`
}

func (req *RequestForShowAuthorizationInfo) Path() string {
	return "/v2/user/auth"
}

func (req *RequestForShowAuthorizationInfo) Method() string {
	return http.MethodGet
}

func (req *RequestForShowAuthorizationInfo) Query() string {
	return ""
}

func (req *RequestForShowAuthorizationInfo) Payload() []byte {
	return nil
}
