package pmbuilder

import (
	"context"
	"net/http"

	"github.com/azrod/go-prometheus-metrics-builder/internal/builder"
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var _ InstanceInterface = &DefaultInstance{}

// DefaultInstance represents the default configuration for the metrics instance.
// It includes options for setting a prefix for metrics and enabling or disabling
// Go runtime metrics collection.
type DefaultInstance struct {
	// PrefixMetric is a string that will be prefixed to all metric names.
	PrefixMetric string

	// DisableGoMetrics is a boolean flag that indicates whether Go runtime metrics
	// should be disabled. If set to true, Go runtime metrics will not be collected.
	DisableGoMetrics bool
}

// InstanceInterface defines the methods required for an instance that can serve HTTP requests,
// handle metrics, and manage configuration settings related to metrics.
type InstanceInterface interface {
	// ListenAndServe starts an HTTP server that serves metrics on the specified address.
	ListenAndServe(ctx context.Context, addr string) error

	// Handler returns an HTTP handler that serves metrics.
	Handler() http.Handler

	// GetPrefixMetric returns the prefix that will be added to all metric names.
	GetPrefixMetric() string

	// SetPrefixMetric sets the prefix that will be added to all metric names.
	SetPrefixMetric(prefix string)

	// GoMetricsIsEnabled returns a boolean flag indicating whether Go runtime metrics are enabled.
	GoMetricsIsEnabled() bool

	// SetGoMetrics sets a boolean flag indicating whether Go runtime metrics should be enabled.
	SetGoMetrics(enabled bool)
}

// New initializes and registers metrics declared in the struct (s)
// The struct must implement the InstanceInterface.
// If a error is found, the function will panic.
func New(s any) {
	x, ok := s.(InstanceInterface)
	if !ok {
		panic("Register require a pointer of pmbuilder.InstanceInterface in the struct")
	}

	builder.BuildMetrics(s, x.GetPrefixMetric(), true)

	if x.GoMetricsIsEnabled() {
		registry.PRegistry.MustRegister(collectors.NewGoCollector())
	}
}

func (i *DefaultInstance) ListenAndServe(ctx context.Context, addr string) error {
	return registry.ListenAndServe(ctx, addr)
}

func (i *DefaultInstance) Handler() http.Handler {
	return registry.Handler()
}

func (i *DefaultInstance) GetPrefixMetric() string {
	if i.PrefixMetric == "" {
		return "pmbuilder"
	}
	return i.PrefixMetric
}

func (i *DefaultInstance) SetPrefixMetric(prefix string) {
	i.PrefixMetric = prefix
}

func (i *DefaultInstance) GoMetricsIsEnabled() bool {
	return !i.DisableGoMetrics
}

func (i *DefaultInstance) SetGoMetrics(goMetrics bool) {
	i.DisableGoMetrics = !goMetrics
}
