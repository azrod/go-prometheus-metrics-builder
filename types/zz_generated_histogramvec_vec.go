package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	HistogramVec struct {
		Metric
		*prometheus.HistogramVec
	}

	histogramVecCollector struct {
		HistogramVec
	}
)

var _ Initializable = &HistogramVec{}

// Init initializes (or re-initializes) the HistogramVec metric
// The metrics is automatically initialized by the Register function
func (c *HistogramVec) Init(mb *Metric) any {
	c.Metric = *mb
	c.HistogramVec = promauto.With(registry.PRegistry).NewHistogramVec(prometheus.HistogramOpts(c.autoBuildHistogramOpts()), mb.Labels)
	return c
}

// Collector returns the prometheus collector
func (c *HistogramVec) Collector() Collector {
	return &histogramVecCollector{
        HistogramVec: *c,
    }
}

// GetType returns the metric type
func (c *HistogramVec) GetType() MetricType {
	return TypeHistogramVec
}

// Run runs the collector.
func (c *histogramVecCollector) SetValue(value float64, labels ...map[string]string) {

	l := prometheus.Labels{}
	bLabels := map[string]string{}
	for _, label := range labels {
		for k, v := range label {
			bLabels[k] = v
		}
	}

	for _, label := range c.Labels {
		if v, ok := bLabels[label]; ok {
			l[label] = v
		}
	}
    c.HistogramVec.With(l).Observe(value)
}

