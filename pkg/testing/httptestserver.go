package testing

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// CreateTestServer creates a http test server using the handler factory, passing in a logger that will pipe the output to the test console
func CreateTestServer(t *testing.T, handlerFactory func(*log.Logger) http.Handler) *httptest.Server {
	logger := testLogger(t)
	handler := handlerFactory(logger)
	server := httptest.NewServer(handler)
	return server
}

func testLogger(t *testing.T) *log.Logger {
	return log.New(testWriter{t}, "test : ", log.LstdFlags)
}

type testWriter struct {
	t *testing.T
}

func (tw testWriter) Write(p []byte) (n int, err error) {
	tw.t.Log(string(p))
	return len(p), nil
}
