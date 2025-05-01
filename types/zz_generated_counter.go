package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	Counter struct {
		Metric
		prometheus.Counter
	}

	counterCollector struct {
		Counter
	}
)

var _ Initializable = &Counter{}

// Init initializes (or re-initializes) the Counter metric
// The metrics is automatically initialized by the Register function
func (c *Counter) Init(mb *Metric) any {
	c.Metric = *mb
	c.Counter = promauto.With(registry.PRegistry).NewCounter(prometheus.CounterOpts(c.autoBuildOpts()))
	return c
}

// Collector returns the prometheus collector
func (c *Counter) Collector() Collector {
	return &counterCollector{
        Counter: *c,
    }
}

// GetType returns the metric type
func (c *Counter) GetType() MetricType {
	return TypeCounter
}

// Run runs the collector.
func (c *counterCollector) SetValue(value float64, _ ...map[string]string) {
	c.Counter.Add(value)
}

