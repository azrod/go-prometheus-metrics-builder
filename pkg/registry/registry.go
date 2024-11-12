package registry

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	PromRegistry interface {
		prometheus.Registerer
		prometheus.Gatherer
	}
)

// PRegistry is the prometheus registry
var PRegistry PromRegistry = prometheus.NewRegistry()

// Handler returns the prometheus handler
func Handler() http.Handler {
	return promhttp.HandlerFor(PRegistry, promhttp.HandlerOpts{
		Registry:      PRegistry,
		ErrorHandling: promhttp.HTTPErrorOnError,
	})
}

// ListenAndServe starts the http server
func ListenAndServe(ctx context.Context, addr string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", Handler())

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		_ = server.Shutdown(context.Background())
	}()

	return server.ListenAndServe()
}
