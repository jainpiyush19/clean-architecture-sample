package adapters

import (
	"context"
	"errors"
	walletdomain "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
	"reflect"
	"testing"
)

func Test_walletImpl_Balance(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func() (context.Context, WalletRepository)
		userID   int64
		want     *walletdomain.Info
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() (context.Context, WalletRepository) {
				mockDb := &dbMock{
					balanceOutput: map[string]interface{}{
						"Amount": 10,
					},
				}
				return createTestRepo(mockDb)
			},
			userID: 1,
			want: &walletdomain.Info{
				Amount:    10,
			},
			wantErr: false,
		},
		{
			name: "error",
			mockFunc: func() (context.Context, WalletRepository) {
				mockDb := &dbMock{err: errors.New("db error")}
				return createTestRepo(mockDb)
			},
			userID:  1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			got, err := w.Balance(ctx, tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("walletImpl.Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walletImpl.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}
