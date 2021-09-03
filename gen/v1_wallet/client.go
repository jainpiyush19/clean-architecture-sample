// Code generated by goa v3.4.3, DO NOT EDIT.
//
// v1/wallet client
//
// Command:
// $ goa gen github.com/jainpiyush19/cryptowallet/design

package v1wallet

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "v1/wallet" service client.
type Client struct {
	HealthEndpoint       goa.Endpoint
	DepositEndpoint      goa.Endpoint
	WithdrawEndpoint     goa.Endpoint
	TransferEndpoint     goa.Endpoint
	BalanceEndpoint      goa.Endpoint
	AdminWalletsEndpoint goa.Endpoint
}

// NewClient initializes a "v1/wallet" service client given the endpoints.
func NewClient(health, deposit, withdraw, transfer, balance, adminWallets goa.Endpoint) *Client {
	return &Client{
		HealthEndpoint:       health,
		DepositEndpoint:      deposit,
		WithdrawEndpoint:     withdraw,
		TransferEndpoint:     transfer,
		BalanceEndpoint:      balance,
		AdminWalletsEndpoint: adminWallets,
	}
}

// Health calls the "health" endpoint of the "v1/wallet" service.
func (c *Client) Health(ctx context.Context) (err error) {
	_, err = c.HealthEndpoint(ctx, nil)
	return
}

// Deposit calls the "deposit" endpoint of the "v1/wallet" service.
func (c *Client) Deposit(ctx context.Context, p *DepositRequestBody) (err error) {
	_, err = c.DepositEndpoint(ctx, p)
	return
}

// Withdraw calls the "withdraw" endpoint of the "v1/wallet" service.
func (c *Client) Withdraw(ctx context.Context, p *WithdrawRequestBody) (err error) {
	_, err = c.WithdrawEndpoint(ctx, p)
	return
}

// Transfer calls the "transfer" endpoint of the "v1/wallet" service.
func (c *Client) Transfer(ctx context.Context, p *TransferRequestBody) (err error) {
	_, err = c.TransferEndpoint(ctx, p)
	return
}

// Balance calls the "balance" endpoint of the "v1/wallet" service.
func (c *Client) Balance(ctx context.Context, p *BalanceRequestBody) (res *BalanceResponseBody, err error) {
	var ires interface{}
	ires, err = c.BalanceEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BalanceResponseBody), nil
}

// AdminWallets calls the "admin/wallets" endpoint of the "v1/wallet" service.
func (c *Client) AdminWallets(ctx context.Context, p *AdminWalletRequestBody) (res *AdminWalletResponseBody, err error) {
	var ires interface{}
	ires, err = c.AdminWalletsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*AdminWalletResponseBody), nil
}