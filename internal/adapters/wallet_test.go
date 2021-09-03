package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dbMock struct {
	dynamodbiface.DynamoDBAPI
	err           error
	balanceOutput map[string]interface{}
}

func (db *dbMock) TransactWriteItemsWithContext(aws.Context, *dynamodb.TransactWriteItemsInput, ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, db.err
}

func (db *dbMock) GetItemWithContext(aws.Context, *dynamodb.GetItemInput, ...request.Option) (*dynamodb.GetItemOutput, error) {
	keys, _ := dynamodbattribute.MarshalMap(db.balanceOutput)
	return &dynamodb.GetItemOutput{Item: keys}, db.err
}

func TestNewWalletRepository(t *testing.T) {
	type args struct {
		dbClient  dynamodbiface.DynamoDBAPI
		tableName string
		historyTable string
	}

	mock := &dbMock{}
	tests := []struct {
		name string
		args args
		want WalletRepository
	}{
		{
			name: "new impl",
			args: args{
				dbClient:  mock,
				tableName: "wallet",
				historyTable: "wallet_history",
			},
			want: &walletImpl{mock, "wallet", "wallet_history"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWalletRepository(tt.args.dbClient, tt.args.tableName, tt.args.historyTable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWalletRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTestRepo(db *dbMock) (context.Context, WalletRepository) {
	ctx := context.Background()
	return ctx, NewWalletRepository(db, "wallet", "wallet_history")
}
