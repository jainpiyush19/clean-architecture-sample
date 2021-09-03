package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	v1wallet "github.com/jainpiyush19/cryptowallet/gen/v1_wallet"
	"github.com/jainpiyush19/cryptowallet/internal/adapters"
	"github.com/jainpiyush19/cryptowallet/internal/app/wallet"
	"github.com/jainpiyush19/cryptowallet/internal/ports"
	awsutil "github.com/jainpiyush19/cryptowallet/pkg/aws"
	httputil "github.com/jainpiyush19/cryptowallet/pkg/http"
	logutil "github.com/jainpiyush19/cryptowallet/pkg/log"
	"github.com/jainpiyush19/cryptowallet/pkg/utils"
)

func main() {
	lgr, flushLogs := logutil.CreateZapLogger()

	awsSession, err := awsutil.CreateAWSSession()
	if err != nil {
		fmt.Println("Got an error creating custom HTTP client:", err)
		return
	}

	// create table in dynamo DB - for demo purposes
	utils.CreateDynamoTables(awsSession)

	walletRepo := adapters.NewWalletRepository(dynamodb.New(awsSession), "wallet", "wallet_history")
	walletUseCase := wallet.NewWalletUseCase(walletRepo)

	// Initialize the services.
	walletSvc := ports.NewWalletService(walletUseCase)
	walletEndpoints := v1wallet.NewEndpoints(walletSvc)

	httputil.StartServer(ports.CreateHTTPHandler(walletEndpoints, lgr), "3000", "", flushLogs)
}
