package coinbasego

import (
	"net/http"
	"time"

	"github.com/sleepychaman/coinbase-go/auth"
)

const ENDPOINT = "https://api.coinbase.com"

type Client struct {
	Auth *auth.Config

	HTTPC       *http.Client
	HTTPTimeout time.Duration
}

func New(auth *auth.Config, fhc *http.Client, timeout time.Duration) *Client {
	if fhc == nil {
		fhc = http.DefaultClient
	}
	return &Client{
		Auth:        auth,
		HTTPC:       fhc,
		HTTPTimeout: timeout,
	}
}
