package coinbasego

import (
	"github.com/sleepychaman/coinbase-go/sign-in-with-coinbase/public"
	"github.com/sleepychaman/coinbase-go/sign-in-with-coinbase/transactions"
)

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

func (p *Client) TransferMoney(req *transactions.RequestForTransferMoney) (*transactions.ResponseForTransferMoney, error) {
	results := new(transactions.ResponseForTransferMoney)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) ListTransactions(req *transactions.RequestForListTransactions) (*transactions.ResponseForListTransactions, error) {
	results := new(transactions.ResponseForListTransactions)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
