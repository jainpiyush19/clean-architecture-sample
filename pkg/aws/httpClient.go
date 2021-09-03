package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

// httpClientSettings defines the HTTP setting for clients
type httpClientSettings struct {
	connect          time.Duration
	connKeepAlive    time.Duration
	expectContinue   time.Duration
	idleConn         time.Duration
	maxAllIdleConns  int
	maxHostIdleConns int
	responseHeader   time.Duration
	tLSHandshake     time.Duration
}

// CreateAWSSession uses the default http client and create an aws session
func CreateAWSSession() (*session.Session, error) {
	httpClient, err := defaultHTTPClient()
	if err != nil {
		fmt.Println("Got an error creating custom HTTP client:", err)
		return nil, err
	}

	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			HTTPClient:       httpClient,
			Credentials:      credentials.NewStaticCredentials("NO-MATTER", "NO-MATTER", ""),
			Endpoint:         aws.String("http://localstack:4566"),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
			Region:           aws.String("us-east-1"),
		},
	}))
	return awsSession, nil
}

// newHTTPClientWithSettings creates an HTTP client with some custom settings
// Inputs:
//     httpSettings contains some custom HTTP settings for the client
// Output:
//     If success, an HTTP client
//     Otherwise, ???
func newHTTPClientWithSettings(httpSettings httpClientSettings) (*http.Client, error) {
	var client http.Client
	tr := &http.Transport{
		ResponseHeaderTimeout: httpSettings.responseHeader,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: httpSettings.connKeepAlive,
			Timeout:   httpSettings.connect,
		}).DialContext,
		MaxIdleConns:          httpSettings.maxAllIdleConns,
		IdleConnTimeout:       httpSettings.idleConn,
		TLSHandshakeTimeout:   httpSettings.tLSHandshake,
		MaxIdleConnsPerHost:   httpSettings.maxHostIdleConns,
		ExpectContinueTimeout: httpSettings.expectContinue,
	}

	// So client makes HTTP/2 requests
	err := http2.ConfigureTransport(tr)
	if err != nil {
		return &client, err
	}

	return &http.Client{
		Transport: tr,
	}, nil
}

//defaultHTTPClient returns default http client tobe used by AWS session
func defaultHTTPClient() (*http.Client, error) {
	return newHTTPClientWithSettings(httpClientSettings{
		connect:          5 * time.Second,
		expectContinue:   1 * time.Second,
		idleConn:         0 * time.Second,
		connKeepAlive:    5 * time.Minute,
		maxAllIdleConns:  0,
		maxHostIdleConns: 500000,
		responseHeader:   5 * time.Second,
		tLSHandshake:     5 * time.Second,
	})
}
