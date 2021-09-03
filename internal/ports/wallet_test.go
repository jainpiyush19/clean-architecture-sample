package ports

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gavv/httpexpect/v2"
	"github.com/jainpiyush19/cryptowallet/internal/domain/user"
	walletdomain "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
	testingtools "github.com/jainpiyush19/cryptowallet/pkg/testing"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"testing"
	"time"

	v1wallet "github.com/jainpiyush19/cryptowallet/gen/v1_wallet"
	"github.com/jainpiyush19/cryptowallet/internal/app/wallet"
)

// TestWalletSvc test for the APIs in wallet
func TestWalletSvc(t *testing.T) {
	testingtools.ExpectTest(t, createTestServer,
		func(e *httpexpect.Expect) {
			e.POST("/deposit").
				WithJSON(&v1wallet.DepositRequestBody{
					Amount: 10,
				}).
				Expect().
				Status(http.StatusCreated)
		},

		func(e *httpexpect.Expect) {
			e.POST("/withdraw").
				WithJSON(&v1wallet.WithdrawRequestBody{
					Amount: 1,
				}).
				Expect().
				Status(http.StatusCreated)
		},

		func(e *httpexpect.Expect) {
			e.POST("/transfer").
				WithJSON(&v1wallet.TransferRequestBody{
					Amount:     1,
					ReceiverID: 2,
				}).
				Expect().
				Status(http.StatusCreated)
		},

		func(e *httpexpect.Expect) {
			e.GET("/balance").
				WithQueryString("userID=1").
				Expect().
				Status(http.StatusOK)
		},

		func(e *httpexpect.Expect) {
			e.GET("/admin/wallets").
				Expect().
				Status(http.StatusOK)
		},
	)
}

//TestPact for contract testing
func TestPact(t *testing.T) {
	//pact.Run(t, "wallet", createTestServer)
}

func TestNewWalletService(t *testing.T) {
	mockUsecase := &wallet.MockUseCases{}

	tests := []struct {
		name          string
		walletUsecase wallet.UseCases
		want          v1wallet.Service
	}{
		{
			name:          "new impl",
			walletUsecase: mockUsecase,
			want:          &walletSvc{mockUsecase},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWalletService(tt.walletUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWalletService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTestService(usecase *wallet.MockUseCases) (context.Context, v1wallet.Service) {
	ctx := context.Background()
	ctx = user.NewContext(ctx, 1)
	return ctx, NewWalletService(usecase)
}

func Test_walletSvc_Health(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func() (context.Context, v1wallet.Service)
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				return createTestService(mockUsecase)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			if err := w.Health(ctx); (err != nil) != tt.wantErr {
				t.Errorf("walletSvc.Health() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_walletSvc_Deposit(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func() (context.Context, v1wallet.Service)
		body     *v1wallet.DepositRequestBody
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Deposit", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return createTestService(mockUsecase)
			},
			body: &v1wallet.DepositRequestBody{
				Amount: 10,
			},
		},
		{
			name: "error",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Deposit", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
				return createTestService(mockUsecase)
			},
			body: &v1wallet.DepositRequestBody{
				Amount: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			if err := w.Deposit(ctx, tt.body); (err != nil) != tt.wantErr {
				t.Errorf("walletSvc.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_walletSvc_Withdraw(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func() (context.Context, v1wallet.Service)
		body     *v1wallet.WithdrawRequestBody
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return createTestService(mockUsecase)
			},
			body: &v1wallet.WithdrawRequestBody{
				Amount: 10,
			},
		},
		{
			name: "error",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
				return createTestService(mockUsecase)
			},
			body: &v1wallet.WithdrawRequestBody{
				Amount: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			if err := w.Withdraw(ctx, tt.body); (err != nil) != tt.wantErr {
				t.Errorf("walletSvc.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_walletSvc_Transfer(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func() (context.Context, v1wallet.Service)
		body     *v1wallet.TransferRequestBody
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Transfer", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return createTestService(mockUsecase)
			},
			body: &v1wallet.TransferRequestBody{
				Amount:     10,
				ReceiverID: 2,
			},
		},
		{
			name: "error",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Transfer", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
				return createTestService(mockUsecase)
			},
			body: &v1wallet.TransferRequestBody{
				Amount:     10,
				ReceiverID: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			if err := w.Transfer(ctx, tt.body); (err != nil) != tt.wantErr {
				t.Errorf("walletSvc.Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_walletSvc_Balance(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		mockFunc func() (context.Context, v1wallet.Service)
		body     *v1wallet.BalanceRequestBody
		wantRes  *v1wallet.BalanceResponseBody
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Balance", mock.Anything, mock.Anything).Return(&walletdomain.Info{
					Amount:    10,
					UpdatedAt: now,
				}, nil)
				return createTestService(mockUsecase)
			},
			body: &v1wallet.BalanceRequestBody{
				UserID: 1,
			},
			wantRes: &v1wallet.BalanceResponseBody{
				Amount:    10,
				UpdatedAt: now.String(),
			},
		},
		{
			name: "error",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("Balance", mock.Anything, mock.Anything).Return(nil, errors.New("error"))
				return createTestService(mockUsecase)
			},
			body: &v1wallet.BalanceRequestBody{
				UserID: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			gotRes, err := w.Balance(ctx, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("walletSvc.Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("walletSvc.Balance() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_walletSvc_AdminWallets(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		mockFunc func() (context.Context, v1wallet.Service)
		body     *v1wallet.AdminWalletRequestBody
		wantRes  *v1wallet.AdminWalletResponseBody
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("AllWallets", mock.Anything).Return([]*walletdomain.Info{&walletdomain.Info{
					UserID:    1,
					Amount:    10,
					UpdatedAt: now,
				}}, nil)
				return createTestService(mockUsecase)
			},
			wantRes: &v1wallet.AdminWalletResponseBody{Data: []*v1wallet.Wallet{{
				UserID:    aws.Int64(1),
				Amount:    aws.Float64(10),
				UpdatedAt: aws.String(now.String()),
			}}},
		},
		{
			name: "error",
			mockFunc: func() (context.Context, v1wallet.Service) {
				mockUsecase := &wallet.MockUseCases{}
				mockUsecase.On("AllWallets", mock.Anything).Return(nil, errors.New("error"))
				return createTestService(mockUsecase)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			gotRes, err := w.AdminWallets(ctx, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("walletSvc.AdminWallets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("walletSvc.AdminWallets() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
