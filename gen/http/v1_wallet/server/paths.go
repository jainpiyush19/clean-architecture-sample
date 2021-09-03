// Code generated by goa v3.4.3, DO NOT EDIT.
//
// HTTP request path constructors for the v1/wallet service.
//
// Command:
// $ goa gen github.com/jainpiyush19/cryptowallet/design

package server

// HealthV1WalletPath returns the URL path to the v1/wallet service health HTTP endpoint.
func HealthV1WalletPath() string {
	return "/health"
}

// DepositV1WalletPath returns the URL path to the v1/wallet service deposit HTTP endpoint.
func DepositV1WalletPath() string {
	return "/deposit"
}

// WithdrawV1WalletPath returns the URL path to the v1/wallet service withdraw HTTP endpoint.
func WithdrawV1WalletPath() string {
	return "/withdraw"
}

// TransferV1WalletPath returns the URL path to the v1/wallet service transfer HTTP endpoint.
func TransferV1WalletPath() string {
	return "/transfer"
}

// BalanceV1WalletPath returns the URL path to the v1/wallet service balance HTTP endpoint.
func BalanceV1WalletPath() string {
	return "/balance"
}

// AdminWalletsV1WalletPath returns the URL path to the v1/wallet service admin/wallets HTTP endpoint.
func AdminWalletsV1WalletPath() string {
	return "/admin/wallets"
}