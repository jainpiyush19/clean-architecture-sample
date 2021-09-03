package adapters

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"time"
)

const Transfer = "Transfer"
const Received = "Received"

//Transfer repository implementation to transfer amount
func (w *walletImpl) Transfer(ctx context.Context, senderID, receiverID int64, amount float64) error {
	if amount < 0 {
		return errors.New("invalid amount. cannot be less than 0")
	}

	senderExp, _ := expression.NewBuilder().
		WithUpdate(
			expression.Add(expression.Name("Amount"), expression.Value(-amount)).
				Set(expression.Name("Date"), expression.Value(time.Now())),
		).
		WithCondition(expression.GreaterThanEqual(expression.Name("Amount"), expression.Value(amount))).
		Build()

	receiverExp, _ := expression.NewBuilder().
		WithUpdate(
			expression.Add(expression.Name("Amount"), expression.Value(amount)).
				Set(expression.Name("Date"), expression.Value(time.Now())),
		).
		Build()

	_, err := w.dbClient.TransactWriteItemsWithContext(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				// add a transaction row to show in history of sender
				Put: &dynamodb.Put{
					Item:      transactionItem(senderID, amount, Transfer, receiverID),
					TableName: aws.String(w.historyTable),
				},
			},
			{
				// add a transaction row to show in history of receiver
				Put: &dynamodb.Put{
					Item:      transactionItem(receiverID, amount, Received, senderID),
					TableName: aws.String(w.historyTable),
				},
			},
			{
				// update balance in sender walletInfo
				Update: &dynamodb.Update{
					ConditionExpression:       senderExp.Condition(),
					ExpressionAttributeNames:  senderExp.Names(),
					ExpressionAttributeValues: senderExp.Values(),
					Key:                       walletDBKeys(senderID),
					TableName:                 aws.String(w.tableName),
					UpdateExpression:          senderExp.Update(),
				},
			},
			{
				// update balance in receiver walletInfo
				Update: &dynamodb.Update{
					ConditionExpression:       receiverExp.Condition(),
					ExpressionAttributeNames:  receiverExp.Names(),
					ExpressionAttributeValues: receiverExp.Values(),
					Key:                       walletDBKeys(receiverID),
					TableName:                 aws.String(w.tableName),
					UpdateExpression:          receiverExp.Update(),
				},
			},
		},
	})
	return err
}
