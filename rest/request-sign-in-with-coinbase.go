package rest

import "github.com/sleepychaman/coinbase-go/rest/sign-in-with-coinbase/public"

func (p *Client) GetPriceSpot(req *public.RequestForGetPriceSpot) (*public.ResponseForGetPriceSpot, error) {
	results := new(public.ResponseForGetPriceSpot)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetPriceBuy(req *public.RequestForGetPriceBuy) (*public.ResponseForGetPriceBuy, error) {
	results := new(public.ResponseForGetPriceBuy)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetPriceSell(req *public.RequestForGetPriceSell) (*public.ResponseForGetPriceSell, error) {
	results := new(public.ResponseForGetPriceSell)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetCurrencies(req *public.RequestForGetCurrencies) (*public.ResponseForGetCurrencies, error) {
	results := new(public.ResponseForGetCurrencies)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
