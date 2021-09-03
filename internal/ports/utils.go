package ports

import (
	"fmt"
	"github.com/jainpiyush19/cryptowallet/gen/http/v1_wallet/server"
	v1wallet "github.com/jainpiyush19/cryptowallet/gen/v1_wallet"
	httputil "github.com/jainpiyush19/cryptowallet/pkg/http"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
	"net/http"
)

// CreateHTTPHandler configures and returns a http.Handler.
func CreateHTTPHandler(walletEndpoints *v1wallet.Endpoints, lgr *zap.Logger) http.Handler {
	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := httputil.JSONResponseEncoder

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	// Setup logger. Replace logger with your own log package of choice.
	eh := httputil.ErrorHandler(lgr)

	walletServer := server.New(walletEndpoints, mux, dec, enc, eh, nil)
	server.Mount(mux, walletServer)
	for _, m := range walletServer.Mounts {
		lgr.Info(fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	return httputil.CreateHTTPHandler(mux, lgr)
}
