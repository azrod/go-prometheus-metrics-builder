package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	Gauge struct {
		Metric
		prometheus.Gauge
	}

	gaugeCollector struct {
		Gauge
	}
)

var _ Initializable = &Gauge{}

// Init initializes (or re-initializes) the Gauge metric
// The metrics is automatically initialized by the Register function
func (c *Gauge) Init(mb *Metric) any {
	c.Metric = *mb
	c.Gauge = promauto.With(registry.PRegistry).NewGauge(prometheus.GaugeOpts(c.autoBuildOpts()))
	return c
}

// Collector returns the prometheus collector
func (c *Gauge) Collector() Collector {
	return &gaugeCollector{
        Gauge: *c,
    }
}

// GetType returns the metric type
func (c *Gauge) GetType() MetricType {
	return TypeGauge
}

// Run runs the collector.
func (c *gaugeCollector) SetValue(value float64, _ ...map[string]string) {
	c.Gauge.Add(value)
    
}
