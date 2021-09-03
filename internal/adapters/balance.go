package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	walletdomain "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
)

//Balance repository implementation to check balance from DB
func (w *walletImpl) Balance(ctx context.Context, userID int64) (*walletdomain.Info, error) {
	out, err := w.dbClient.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		Key:       walletDBKeys(userID),
		TableName: aws.String(w.tableName),
	})

	if err != nil {
		return nil, err
	}

	wallet := &walletInfo{}
	err = dynamodbattribute.UnmarshalMap(out.Item, wallet)
	if err != nil {
		return nil, err
	}

	return &walletdomain.Info{
		Amount:    wallet.Amount,
		UpdatedAt: wallet.Date,
	}, nil
}
