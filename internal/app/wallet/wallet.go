package wallet

import (
	"context"
	"errors"
	"github.com/jainpiyush19/cryptowallet/internal/adapters"
	walletdomain "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
)

//UseCases for wallet
type UseCases interface {
	Deposit(ctx context.Context, userID int64, amount float64) error
	Withdraw(ctx context.Context, userID int64, amount float64) error
	Transfer(ctx context.Context, senderID, receiverID int64, amount float64) error
	Balance(ctx context.Context, userID int64) (*walletdomain.Info, error)
	AllWallets(ctx context.Context) ([]*walletdomain.Info, error)
}

//impl implementation of wallet usecase
type impl struct {
	repo adapters.WalletRepository
}

//NewWalletUseCase returns a new wallet usecase implementation
func NewWalletUseCase(repo adapters.WalletRepository) UseCases {
	return &impl{repo}
}

//Deposit funds in wallet
func (i *impl) Deposit(ctx context.Context, userID int64, amount float64) error {
	if userID < 1 {
		return errors.New("dsadas")
	}
	return i.repo.Deposit(ctx, userID, amount)
}

//Withdraw funds from wallet
func (i *impl) Withdraw(ctx context.Context, userID int64, amount float64) error {
	return i.repo.Withdraw(ctx, userID, amount)
}

//Transfer funds from wallet to another user
func (i *impl) Transfer(ctx context.Context, senderID, receiverID int64, amount float64) error {
	return i.repo.Transfer(ctx, senderID, receiverID, amount)
}

//Balance returns funds in wallet
func (i *impl) Balance(ctx context.Context, userID int64) (*walletdomain.Info, error) {
	return i.repo.Balance(ctx, userID)
}

//AllWallets returns all wallets
func (i *impl) AllWallets(ctx context.Context) ([]*walletdomain.Info, error) {
	return i.repo.AllWallets(ctx)
}
