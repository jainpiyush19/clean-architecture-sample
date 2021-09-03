package adapters

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"time"
)

const Deposit = "Deposit"

//Deposit repository implementation to deposit amount in walletInfo
func (w *walletImpl) Deposit(ctx context.Context, userID int64, amount float64) error {
	if amount < 0 {
		return errors.New("invalid amount. cannot be less than 0")
	}

	// to update amount in DB
	exp, _ := expression.NewBuilder().
		WithUpdate(
			expression.Add(expression.Name("Amount"), expression.Value(amount)).
				Set(expression.Name("Date"), expression.Value(time.Now())),
		).
		Build()

	_, err := w.dbClient.TransactWriteItemsWithContext(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				// add a transaction row to show in history
				Put: &dynamodb.Put{
					Item:      transactionItem(userID, amount, Deposit, 0),
					TableName: aws.String(w.historyTable),
				},
			},
			{
				// update the balance
				Update: &dynamodb.Update{
					ConditionExpression:       exp.Condition(),
					ExpressionAttributeNames:  exp.Names(),
					ExpressionAttributeValues: exp.Values(),
					Key:                       walletDBKeys(userID),
					TableName:                 aws.String(w.tableName),
					UpdateExpression:          exp.Update(),
				},
			},
		},
	})
	return err
}
