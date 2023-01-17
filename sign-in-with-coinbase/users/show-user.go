package users

import (
	"net/http"
)

type RequestForShowUser struct {
	UserID string
}

type ResponseForShowUser struct {
	Data User `json:"data"`
}

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileLocation string `json:"profile_location"`
	ProfileBio      string `json:"profile_bio"`
	ProfileUrl      string `json:"profile_url"`
	AvatarUrl       string `json:"avatar_url"`
	Ressource       string `json:"resource"`
	RessourcePath   string `json:"resource_path"`
}

func (req *RequestForShowUser) Path() string {
	return "/v2/users/" + req.UserID
}

func (req *RequestForShowUser) Method() string {
	return http.MethodGet
}

func (req *RequestForShowUser) Query() string {
	return ""
}

func (req *RequestForShowUser) Payload() []byte {
	return nil
}
