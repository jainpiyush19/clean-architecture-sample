// Code generated by goa v3.4.3, DO NOT EDIT.
//
// v1/wallet service
//
// Command:
// $ goa gen github.com/jainpiyush19/cryptowallet/design

package v1wallet

import (
	"context"

	"goa.design/goa/v3/security"
)

// wallet service contains APIs for wallet transactions
type Service interface {
	// This API checks for status 200 for downstream services
	Health(context.Context) (err error)
	// This API deposit money in wallet
	Deposit(context.Context, *DepositRequestBody) (err error)
	// This API withdraw money from wallet
	Withdraw(context.Context, *WithdrawRequestBody) (err error)
	// This API transfer money from one wallet to another
	Transfer(context.Context, *TransferRequestBody) (err error)
	// This API checks balance in wallet
	Balance(context.Context, *BalanceRequestBody) (res *BalanceResponseBody, err error)
	// This API returns all wallets
	AdminWallets(context.Context, *AdminWalletRequestBody) (res *AdminWalletResponseBody, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "v1/wallet"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"health", "deposit", "withdraw", "transfer", "balance", "admin/wallets"}

// DepositRequestBody is the payload type of the v1/wallet service deposit
// method.
type DepositRequestBody struct {
	// JWT token used to perform authorization
	Token  *string
	Amount float64
}

// WithdrawRequestBody is the payload type of the v1/wallet service withdraw
// method.
type WithdrawRequestBody struct {
	// JWT token used to perform authorization
	Token  *string
	Amount float64
}

// TransferRequestBody is the payload type of the v1/wallet service transfer
// method.
type TransferRequestBody struct {
	// JWT token used to perform authorization
	Token      *string
	ReceiverID int64
	Amount     float64
}

// BalanceRequestBody is the payload type of the v1/wallet service balance
// method.
type BalanceRequestBody struct {
	// JWT token used to perform authorization
	Token *string
	// user ID
	UserID int64
}

// BalanceResponseBody is the result type of the v1/wallet service balance
// method.
type BalanceResponseBody struct {
	Amount    float64
	UpdatedAt string
}

// AdminWalletRequestBody is the payload type of the v1/wallet service
// admin/wallets method.
type AdminWalletRequestBody struct {
	// JWT token used to perform authorization
	Token *string
}

// AdminWalletResponseBody is the result type of the v1/wallet service
// admin/wallets method.
type AdminWalletResponseBody struct {
	Data []*Wallet
}

type Wallet struct {
	UserID    *int64
	Amount    *float64
	UpdatedAt *string
}
