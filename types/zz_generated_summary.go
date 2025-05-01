package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	Summary struct {
		Metric
		prometheus.Summary
	}

	summaryCollector struct {
		Summary
	}
)

var _ Initializable = &Summary{}

// Init initializes (or re-initializes) the Summary metric
// The metrics is automatically initialized by the Register function
func (c *Summary) Init(mb *Metric) any {
	c.Metric = *mb
    c.Summary = promauto.With(registry.PRegistry).NewSummary(prometheus.SummaryOpts(c.autoBuildSummaryOpts()))
	return c
}

// Collector returns the prometheus collector
func (c *Summary) Collector() Collector {
	return &summaryCollector{
        Summary: *c,
    }
}

// GetType returns the metric type
func (c *Summary) GetType() MetricType {
	return TypeSummary
}

// Run runs the collector.
func (c *summaryCollector) SetValue(value float64, _ ...map[string]string) {
    c.Summary.Observe(value)
}

