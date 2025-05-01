package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	CounterVec struct {
		Metric
		*prometheus.CounterVec
	}

	counterVecCollector struct {
		CounterVec
	}
)

var _ Initializable = &CounterVec{}

// Init initializes (or re-initializes) the CounterVec metric
// The metrics is automatically initialized by the Register function
func (c *CounterVec) Init(mb *Metric) any {
	c.Metric = *mb
	c.CounterVec = promauto.With(registry.PRegistry).NewCounterVec(prometheus.CounterOpts(c.autoBuildOpts()), mb.Labels)
	return c
}

// Collector returns the prometheus collector
func (c *CounterVec) Collector() Collector {
	return &counterVecCollector{
		CounterVec: *c,
	}
}

// GetType returns the metric type
func (c *CounterVec) GetType() MetricType {
	return TypeCounterVec
}

// Run runs the collector.
func (c *counterVecCollector) SetValue(value float64, labels ...map[string]string) {

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
	c.CounterVec.With(l).Add(value)
}
