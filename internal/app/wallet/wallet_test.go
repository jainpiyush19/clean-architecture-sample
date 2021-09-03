package wallet

import (
	"context"
	"errors"
	"github.com/jainpiyush19/cryptowallet/internal/adapters"
	walletdomain "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

func TestNewWalletUseCase(t *testing.T) {
	mockRepo := &adapters.MockWalletRepository{}
	tests := []struct {
		name string
		repo adapters.WalletRepository
		want UseCases
	}{
		{
			name: "new impl",
			repo: mockRepo,
			want: &impl{mockRepo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWalletUseCase(tt.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWalletUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

//createTestUseCase creates a test usecase impl with mockrepo
func createTestUseCase(repo *adapters.MockWalletRepository) (context.Context, UseCases) {
	ctx := context.Background()
	return ctx, NewWalletUseCase(repo)
}

type args struct {
	userID int64
	amount float64
}

var testArgs = args{
	userID: 1,
	amount: 10,
}

func Test_impl_Deposit(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func() (context.Context, UseCases)
		args     args
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("Deposit", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return createTestUseCase(mockRepo)
			},
			args:    testArgs,
			wantErr: false,
		},
		{
			name: "error",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("Deposit", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("db error"))
				return createTestUseCase(mockRepo)
			},
			args:    testArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, impl := tt.mockFunc()
			if err := impl.Deposit(ctx, tt.args.userID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("impl.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func Test_impl_Withdraw(t *testing.T) {
//	tests := []struct {
//		name     string
//		mockFunc func() (context.Context, UseCases)
//		args     args
//		wantErr  bool
//	}{
//		{
//			name: "success",
//			mockFunc: func() (context.Context, UseCases) {
//				mockRepo := &adapters.MockWalletRepository{}
//				mockRepo.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(nil)
//				return createTestUseCase(mockRepo)
//			},
//			args:    testArgs,
//			wantErr: false,
//		},
//		{
//			name: "error",
//			mockFunc: func() (context.Context, UseCases) {
//				mockRepo := &adapters.MockWalletRepository{}
//				mockRepo.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("db error"))
//				return createTestUseCase(mockRepo)
//			},
//			args:    testArgs,
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ctx, impl := tt.mockFunc()
//			if err := impl.Withdraw(ctx, tt.args.userID, tt.args.amount); (err != nil) != tt.wantErr {
//				t.Errorf("impl.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func Test_impl_Transfer(t *testing.T) {
	type args struct {
		senderID   int64
		receiverID int64
		amount     float64
	}

	tests := []struct {
		name     string
		mockFunc func() (context.Context, UseCases)
		args     args
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("Transfer", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return createTestUseCase(mockRepo)
			},
			args: args{
				senderID:   1,
				receiverID: 2,
				amount:     10,
			},
			wantErr: false,
		},
		{
			name: "error",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("Transfer", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("db error"))
				return createTestUseCase(mockRepo)
			},
			args: args{
				senderID:   1,
				receiverID: 2,
				amount:     10,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, impl := tt.mockFunc()
			if err := impl.Transfer(ctx, tt.args.senderID, tt.args.receiverID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("impl.Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_impl_Balance(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		mockFunc func() (context.Context, UseCases)
		userID   int64
		want     *walletdomain.Info
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("Balance", mock.Anything, mock.Anything).Return(&walletdomain.Info{
					Amount:    10,
					UpdatedAt: now,
				}, nil)
				return createTestUseCase(mockRepo)
			},
			userID: 1,
			want: &walletdomain.Info{
				Amount:    10,
				UpdatedAt: now,
			},
			wantErr: false,
		},
		{
			name: "error",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("Balance", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
				return createTestUseCase(mockRepo)
			},
			userID:  1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, impl := tt.mockFunc()
			got, err := impl.Balance(ctx, tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_AllWallets(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		mockFunc func() (context.Context, UseCases)
		want     []*walletdomain.Info
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("AllWallets", mock.Anything).Return([]*walletdomain.Info{{
					UserID:    1,
					Amount:    10,
					UpdatedAt: now,
				}}, nil)
				return createTestUseCase(mockRepo)
			},
			want: []*walletdomain.Info{{
				UserID:    1,
				Amount:    10,
				UpdatedAt: now,
			}},
			wantErr: false,
		},
		{
			name: "error",
			mockFunc: func() (context.Context, UseCases) {
				mockRepo := &adapters.MockWalletRepository{}
				mockRepo.On("AllWallets", mock.Anything).Return(nil, errors.New("db error"))
				return createTestUseCase(mockRepo)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, impl := tt.mockFunc()
			got, err := impl.AllWallets(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.AllWallets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.AllWallets() = %v, want %v", got, tt.want)
			}
		})
	}
}
