package http

import (
	"context"
	"fmt"
	"github.com/NYTimes/gziphandler"
	logutil "github.com/jainpiyush19/cryptowallet/pkg/log"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
	httpmiddleware "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CreateHTTPHandler configures and returns a http.Handler
func CreateHTTPHandler(mux goahttp.Muxer, logger *zap.Logger) http.Handler {
	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	handler = gziphandler.GzipHandler(handler)
	handler = httpmiddleware.Log(logutil.NewLogger(logger))(handler)
	handler = httpmiddleware.RequestID(middleware.UseRequestIDOption(true))(handler)

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	return handler
}

//StartServer starts a new http server
func StartServer(h http.Handler, defaultPort string, env string, flushLogs func()) {
	log.Println("Starting http listener")

	if defaultPort == "" {
		defaultPort = "3000"
	}
	addr := ":" + defaultPort
	if env == "local" {
		addr = "localhost" + addr
	}

	server := &http.Server{Addr: addr, Handler: h}
	go func() {
		log.Println("Starting server on " + addr)
		log.Fatal(server.ListenAndServe())
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Waiting for SIGINT (pkill -2)
	<-stop
	log.Println("Shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// flushes any logs in the buffer
	flushLogs()
	if err := server.Shutdown(ctx); err != nil {
		// handle err
	}
}

//JSONResponseEncoder uses jsoniter to encode response with content-type = application/json
func JSONResponseEncoder(ctx context.Context, w http.ResponseWriter) goahttp.Encoder {
	goahttp.SetContentType(w, "application/json")
	return jsoniter.NewEncoder(w)
}

// ErrorHandler returns a function that writes and logs the given errors.
// The function also writes and logs the errors unique ID so that it's possible
// to correlate.
func ErrorHandler(logger *zap.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		logger.Error(fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
