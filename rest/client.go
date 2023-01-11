package rest

import (
	"time"

	"github.com/sleepychaman/coinbase-go/auth"
	"github.com/valyala/fasthttp"
)

const ENDPOINT = "https://api.coinbase.com"

type Client struct {
	Auth *auth.Config

	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func New(auth *auth.Config, fhc *fasthttp.Client, timeout time.Duration) *Client {
	if fhc == nil {
		fhc = new(fasthttp.Client)
	}
	return &Client{
		Auth:        auth,
		HTTPC:       fhc,
		HTTPTimeout: timeout,
	}
}
