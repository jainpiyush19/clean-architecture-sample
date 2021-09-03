package ports

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	v1wallet "github.com/jainpiyush19/cryptowallet/gen/v1_wallet"
	"github.com/jainpiyush19/cryptowallet/internal/adapters"
	"github.com/jainpiyush19/cryptowallet/internal/app/wallet"
	logutil "github.com/jainpiyush19/cryptowallet/pkg/log"
	testingtools "github.com/jainpiyush19/cryptowallet/pkg/testing"
	"github.com/jainpiyush19/cryptowallet/pkg/testing/testcontainers/localstack"
	utils2 "github.com/jainpiyush19/cryptowallet/pkg/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var awsSession *session.Session

func TestMain(m *testing.M) {
	log.Println("Start Test Containers")
	ctx := context.Background()

	c, err := localstack.NewContainerBuilder().
		WithServices(localstack.DynamoDB).
		Build()
	if err != nil {
		log.Fatal("Could not start test containers - ", err)
	}
	endpoint, _ := c.Endpoint()
	awsSession, _ = session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("NO-MATTER", "NO-MATTER", ""),
		Endpoint:         aws.String(endpoint),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("us-east-1"),
	})
	utils2.CreateDynamoTables(awsSession)

	exitVal := m.Run()
	c.Terminate(ctx)
	os.Exit(exitVal)
}

func createTestServer(t *testing.T) *httptest.Server {
	return testingtools.CreateTestServer(t, func(logger *log.Logger) http.Handler {
		walletRepo := adapters.NewWalletRepository(dynamodb.New(awsSession), "wallet", "wallet_history")
		walletUseCase := wallet.NewWalletUseCase(walletRepo)

		// Initialize the services.
		walletSvc := NewWalletService(walletUseCase)
		walletEndpoints := v1wallet.NewEndpoints(walletSvc)

		lgr, flushLogs := logutil.CreateZapLogger()
		defer flushLogs()

		return CreateHTTPHandler(walletEndpoints, lgr)
	})
}
