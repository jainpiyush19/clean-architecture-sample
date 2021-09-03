package testing

import (
	"fmt"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

// RunK6LoadTest allows you to execute load tests using k6. k6 must be installed to use this
func RunK6LoadTest(t *testing.T, testServerFactory func(t *testing.T) *httptest.Server, scriptFile string) {
	server := testServerFactory(t)
	defer server.Close()

	cmd := exec.Command("k6", "run", "-e", fmt.Sprintf("BASE_URL=%s", server.URL), scriptFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}
}
