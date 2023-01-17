package coinbasego

import (
	"github.com/sleepychaman/coinbase-go/advanded-trade-api/account"
	"github.com/sleepychaman/coinbase-go/advanded-trade-api/order"
	"github.com/sleepychaman/coinbase-go/advanded-trade-api/products"
)

func (p *Client) ListAccounts(req *account.RequestForListAccounts) (*account.ResponseForListAccounts, error) {
	results := new(account.ResponseForListAccounts)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetAccount(req *account.RequestForGetAccount) (*account.ResponseForGetAccount, error) {
	results := new(account.ResponseForGetAccount)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CreateOrder(req *order.RequestForCreateOrder) (*order.ResponseForCreateOrder, error) {
	results := new(order.ResponseForCreateOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CancelOrders(req *order.RequestForCancelOrders) (*order.ResponseForCancelOrders, error) {
	results := new(order.ResponseForCancelOrders)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetOrder(req *order.RequestForGetOrder) (*order.ResponseForGetOrder, error) {
	results := new(order.ResponseForGetOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetProduct(req *products.RequestForGetProduct) (*products.ResponseForGetProduct, error) {
	results := new(products.ResponseForGetProduct)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
