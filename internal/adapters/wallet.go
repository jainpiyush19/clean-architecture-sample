package adapters

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	walletdomain "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
	"time"
)

//WalletRepository provides methods available on walletInfo repository like deposit, transfer, withdraw, balance
type WalletRepository interface {
	Deposit(ctx context.Context, userID int64, amount float64) error
	Withdraw(ctx context.Context, userID int64, amount float64) error
	Transfer(ctx context.Context, senderID, receiverID int64, amount float64) error
	Balance(ctx context.Context, userID int64) (*walletdomain.Info, error)
	AllWallets(ctx context.Context) ([]*walletdomain.Info, error)
}

//walletImpl of repository
type walletImpl struct {
	dbClient     dynamodbiface.DynamoDBAPI
	tableName    string
	historyTable string
}

//NewWalletRepository returns a new implementation of walletInfo repository
func NewWalletRepository(dbClient dynamodbiface.DynamoDBAPI, tableName, historyTable string) WalletRepository {
	return &walletImpl{dbClient, tableName, historyTable}
}

//transaction info for a transaction
type transaction struct {
	PartitionKey    string
	RangeKey        string
	Amount          float64
	Date            time.Time
	transactionType string
	receiverID      int64 // in case of transfers
}

//walletInfo info for walletInfo
type walletInfo struct {
	PartitionKey string
	RangeKey     string
	Amount       float64
	Date         time.Time
}

//walletDBKeys return primary key for DB
func walletDBKeys(userID int64) map[string]*dynamodb.AttributeValue {
	keys := map[string]string{
		"PartitionKey": fmt.Sprintf("USER#%d", userID),
		"RangeKey":     "Wallet",
	}

	dbKeys, _ := dynamodbattribute.MarshalMap(keys)
	return dbKeys
}

//transactionItem return transaction row for DB
func transactionItem(userID int64, amount float64, transactionType string, receiverID int64) map[string]*dynamodb.AttributeValue {
	now := time.Now()
	item := &transaction{
		PartitionKey:    fmt.Sprintf("USER#%d", userID),
		RangeKey:        fmt.Sprintf("Transaction#%s", now),
		Amount:          amount,
		Date:            now,
		transactionType: transactionType,
		receiverID:      receiverID,
	}

	dbItem, _ := dynamodbattribute.MarshalMap(item)
	return dbItem
}
