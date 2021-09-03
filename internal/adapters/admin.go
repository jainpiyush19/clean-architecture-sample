package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	walletdomain "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
	"strconv"
	"strings"
)

//AllWallets repository implementation to return all wallets
func (w *walletImpl) AllWallets(ctx context.Context) ([]*walletdomain.Info, error) {
	allWallets := []*walletdomain.Info{}
	err := w.dbClient.ScanPagesWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(w.tableName),
	}, func(output *dynamodb.ScanOutput, b bool) bool {
		wallets := []*walletInfo{}
		err := dynamodbattribute.UnmarshalListOfMaps(output.Items, &wallets)
		if err != nil {
			return false
		}

		for _, wallet := range wallets {
			userID, _ := strconv.Atoi(strings.Split(wallet.PartitionKey, "#")[1])
			allWallets = append(allWallets, &walletdomain.Info{
				UserID:    int64(userID),
				Amount:    wallet.Amount,
				UpdatedAt: wallet.Date,
			})
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	return allWallets, nil
}
