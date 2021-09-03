package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"time"
)

const Withdraw = "Withdraw"

//Withdraw repository implementation to withdraw amount
func (w *walletImpl) Withdraw(ctx context.Context, userID int64, amount float64) error {
	exp, _ := expression.NewBuilder().
		WithUpdate(
			expression.Add(expression.Name("Amount"), expression.Value(-amount)).
				Set(expression.Name("Date"), expression.Value(time.Now())),
		).
		WithCondition(expression.GreaterThanEqual(expression.Name("Amount"), expression.Value(amount))).
		Build()

	_, err := w.dbClient.TransactWriteItemsWithContext(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				// add a transaction row to show in history
				Put: &dynamodb.Put{
					Item:      transactionItem(userID, amount, Withdraw, 0),
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
