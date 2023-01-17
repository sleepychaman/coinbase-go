package coinbasego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/sleepychaman/coinbase-go/generic"
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

func (p *Client) newRequest(r Requester) *http.Request {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(ENDPOINT)
	u.Path = u.Path + r.Path()
	u.RawQuery = r.Query()
	reader := bytes.NewReader(r.Payload())
	req, _ := http.NewRequest(r.Method(), u.String(), reader)
	body := r.Payload()

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
		req.Header.Set("CB-VERSION", "2023-01-01")
		req.Header.Set("CB-ACCESS-KEY", p.Auth.Key)
		req.Header.Set("CB-ACCESS-SIGN", p.Auth.Signature(message))
		req.Header.Set("CB-ACCESS-TIMESTAMP", timestamp)
	}
	return req
}

func (c *Client) do(r Requester) (*http.Response, error) {
	req := c.newRequest(r)
	fmt.Println(req)
	res, err := c.HTTPC.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		var r interface{}
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, &generic.APIError{
				Status:  res.StatusCode,
				Message: fmt.Sprintf("%+v\n", string(data)),
			}
		}
		if resp, ok := r.(generic.APIError); ok {
			return nil, resp
		}
		return nil, &generic.APIError{
			Status:  res.StatusCode,
			Message: fmt.Sprintf("%+v\n", string(data)),
		}

	}
	return res, nil
}

func decode(res *http.Response, out interface{}) error {
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, out); err != nil {
		return err
	}
	return nil
}
