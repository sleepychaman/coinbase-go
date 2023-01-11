package rest

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/sleepychaman/coinbase-go/rest/generic"
	"github.com/valyala/fasthttp"
)

func (p *Client) request(req Requester, results interface{}) error {
	res, err := p.do(req)
	if err != nil {
		return err
	}

	if err := decode(res, results); err != nil {
		return err
	}
	return nil
}

func (p *Client) newRequest(r Requester) *fasthttp.Request {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(ENDPOINT)
	u.Path = u.Path + r.Path()
	u.RawQuery = r.Query()

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(u.String())
	body := r.Payload()
	req.SetBody(body)

	if p.Auth != nil {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		message := fmt.Sprintf(
			"%s%s%s%s",
			timestamp,
			r.Method(),
			r.Path(),
			body,
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("CB-ACCESS-KEY", p.Auth.Key)
		req.Header.Set("CB-ACCESS-SIGN", p.Auth.Signature(message))
		req.Header.Set("CB-ACCESS-TIMESTAMP", timestamp)
	}
	return req
}

func (c *Client) do(r Requester) (*fasthttp.Response, error) {
	req := c.newRequest(r)
	fmt.Println(req)
	// fasthttp for http2.0
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", string(res.Body()))
	// no usefull headers

	if res.StatusCode() != 200 {
		var r interface{}
		if err := json.Unmarshal(res.Body(), &r); err != nil {
			return nil, &generic.APIError{
				Status:  res.StatusCode(),
				Message: fmt.Sprintf("%+v\n", string(res.Body())),
			}
		}
		if resp, ok := r.(generic.APIError); ok {
			return nil, resp
		}
		return nil, &generic.APIError{
			Status:  res.StatusCode(),
			Message: fmt.Sprintf("%+v\n", string(res.Body())),
		}

	}
	return res, nil
}

func decode(res *fasthttp.Response, out interface{}) error {
	if err := json.Unmarshal(res.Body(), out); err != nil {
		return err
	}
	return nil
}
