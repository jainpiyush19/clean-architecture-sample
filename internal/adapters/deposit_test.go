package adapters

import (
	"context"
	"errors"
	"testing"
)

func Test_walletImpl_Deposit(t *testing.T) {
	type args struct {
		userID int64
		amount float64
	}

	tests := []struct {
		name     string
		mockFunc func() (context.Context, WalletRepository)
		args     args
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, WalletRepository) {
				mockDb := &dbMock{}
				return createTestRepo(mockDb)
			},
			args: args{
				userID: 1,
				amount: 10,
			},
			wantErr: false,
		},
		{
			name: "error",
			mockFunc: func() (context.Context, WalletRepository) {
				mockDb := &dbMock{err: errors.New("db error")}
				return createTestRepo(mockDb)
			},
			args: args{
				userID: 1,
				amount: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			if err := w.Deposit(ctx, tt.args.userID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("walletImpl.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
