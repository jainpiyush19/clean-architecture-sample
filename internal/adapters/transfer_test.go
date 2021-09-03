package adapters

import (
	"context"
	"errors"
	"testing"
)

func Test_walletImpl_Transfer(t *testing.T) {
	type args struct {
		senderID   int64
		receiverID int64
		amount     float64
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
				senderID:   1,
				receiverID: 2,
				amount:     10,
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
				senderID:   1,
				receiverID: 2,
				amount:     10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := tt.mockFunc()
			if err := w.Transfer(ctx, tt.args.senderID, tt.args.receiverID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("walletImpl.Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
