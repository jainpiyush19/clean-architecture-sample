package testing

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

// ExpectTest allows you to run use httpexpect APIs while managing the http test server lifecycle for you
func ExpectTest(t *testing.T, testServerFactory func(t *testing.T) *httptest.Server, test ...func(httpexpect *httpexpect.Expect)) {
	server := testServerFactory(t)
	defer server.Close()
	expect := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	for _, t := range test {
		t(expect)
	}
}
