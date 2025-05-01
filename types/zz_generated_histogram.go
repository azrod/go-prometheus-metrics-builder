package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	Histogram struct {
		Metric
		prometheus.Histogram
	}

	histogramCollector struct {
		Histogram
	}
)

var _ Initializable = &Histogram{}

// Init initializes (or re-initializes) the Histogram metric
// The metrics is automatically initialized by the Register function
func (c *Histogram) Init(mb *Metric) any {
	c.Metric = *mb
	c.Histogram = promauto.With(registry.PRegistry).NewHistogram(prometheus.HistogramOpts(c.autoBuildHistogramOpts()))
	return c
}

// Collector returns the prometheus collector
func (c *Histogram) Collector() Collector {
	return &histogramCollector{
        Histogram: *c,
    }
}

// GetType returns the metric type
func (c *Histogram) GetType() MetricType {
	return TypeHistogram
}

// Run runs the collector.
func (c *histogramCollector) SetValue(value float64, _ ...map[string]string) {
    c.Histogram.Observe(value)
}

func (c Histogram) NewTimer() *prometheus.Timer {
	return prometheus.NewTimer(c.Histogram)
}
