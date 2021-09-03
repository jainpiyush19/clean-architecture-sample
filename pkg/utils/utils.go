package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

//CreateDynamoTables creates a new dynamo db table for wallet
func CreateDynamoTables(sess *session.Session) {
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	fmt.Println("Creating the tables")

	tableName := "wallet"
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("PartitionKey"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("RangeKey"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("PartitionKey"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("RangeKey"),
				KeyType:       aws.String("RANGE"),
			},
		},
		BillingMode: aws.String("PAY_PER_REQUEST"),
		TableName:   aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println("Created the table")
			}
		} else {
			fmt.Println("Got error calling CreateTable:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	input.TableName = aws.String("wallet_history")
	_, err = svc.CreateTable(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println("Created the table")
			}
		} else {
			fmt.Println("Got error calling CreateTable:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	fmt.Println("Created the tables")
}
