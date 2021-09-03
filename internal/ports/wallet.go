package ports

import (
	"context"
	"errors"
	v1wallet "github.com/jainpiyush19/cryptowallet/gen/v1_wallet"
	"github.com/jainpiyush19/cryptowallet/internal/app/wallet"
	"github.com/jainpiyush19/cryptowallet/internal/domain/user"
	"goa.design/goa/v3/security"
)

//walletSvc implementation
type walletSvc struct {
	wallet wallet.UseCases
}

//NewWalletService to initialize new service object
func NewWalletService(walletUsecase wallet.UseCases) v1wallet.Service {
	return &walletSvc{walletUsecase}
}

// JWTAuth implements the authorization logic for service "yoda" for the "jwt"
// security scheme.
func (w *walletSvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	// JWT validation skipped for now
	ctx = user.NewContext(ctx, 1)
	return ctx, nil
}

//fetchUserFromContext fetch user from context
func fetchUserFromContext(ctx context.Context) int64 {
	userID, _ := user.FromContext(ctx)
	if userID == nil {
		return 0
	}
	return *userID
}

//Health check API handler
func (w *walletSvc) Health(ctx context.Context) (err error) {
	return nil
}

//Deposit API handler
func (w *walletSvc) Deposit(ctx context.Context, body *v1wallet.DepositRequestBody) (err error) {
	userID := fetchUserFromContext(ctx)
	if userID < 1 {
		return errors.New(InvalidUser)
	}
	return w.wallet.Deposit(ctx, userID, body.Amount)
}

//Withdraw API handler
func (w *walletSvc) Withdraw(ctx context.Context, body *v1wallet.WithdrawRequestBody) (err error) {
	userID := fetchUserFromContext(ctx)
	if userID < 1 {
		return errors.New(InvalidUser)
	}
	return w.wallet.Withdraw(ctx, userID, body.Amount)
}

//Transfer API handler
func (w *walletSvc) Transfer(ctx context.Context, body *v1wallet.TransferRequestBody) (err error) {
	userID := fetchUserFromContext(ctx)
	if userID < 1 {
		return errors.New(InvalidUser)
	}

	if userID == body.ReceiverID {
		return errors.New(InvalidReceiver)
	}

	return w.wallet.Transfer(ctx, userID, body.ReceiverID, body.Amount)
}

//Balance API handler
func (w *walletSvc) Balance(ctx context.Context, body *v1wallet.BalanceRequestBody) (res *v1wallet.BalanceResponseBody, err error) {
	//userID := fetchUserFromContext(ctx)
	userID := body.UserID
	if userID < 1 {
		return nil, errors.New(InvalidUser)
	}

	balance, err := w.wallet.Balance(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &v1wallet.BalanceResponseBody{Amount: balance.Amount, UpdatedAt: balance.UpdatedAt.String()}, nil
}

//AdminWallets API handler
func (w *walletSvc) AdminWallets(ctx context.Context, body *v1wallet.AdminWalletRequestBody) (res *v1wallet.AdminWalletResponseBody, err error) {
	//check for admin priveleges
	wallets, err := w.wallet.AllWallets(ctx)
	if err != nil {
		return nil, err
	}

	resWallets := []*v1wallet.Wallet{}
	for _, info := range wallets {
		updatedAt := info.UpdatedAt.String()
		resWallets = append(resWallets, &v1wallet.Wallet{
			UserID:    &info.UserID,
			Amount:    &info.Amount,
			UpdatedAt: &updatedAt,
		})
	}

	res = &v1wallet.AdminWalletResponseBody{Data: resWallets}
	return res, err
}
