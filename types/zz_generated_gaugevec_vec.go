package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	GaugeVec struct {
		Metric
		*prometheus.GaugeVec
	}

	gaugeVecCollector struct {
		GaugeVec
	}
)

var _ Initializable = &GaugeVec{}

// Init initializes (or re-initializes) the GaugeVec metric
// The metrics is automatically initialized by the Register function
func (c *GaugeVec) Init(mb *Metric) any {
	c.Metric = *mb
	c.GaugeVec = promauto.With(registry.PRegistry).NewGaugeVec(prometheus.GaugeOpts(c.autoBuildOpts()), mb.Labels)
	return c
}

// Collector returns the prometheus collector
func (c *GaugeVec) Collector() Collector {
	return &gaugeVecCollector{
        GaugeVec: *c,
    }
}

// GetType returns the metric type
func (c *GaugeVec) GetType() MetricType {
	return TypeGaugeVec
}

// Run runs the collector.
func (c *gaugeVecCollector) SetValue(value float64, labels ...map[string]string) {

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
	c.GaugeVec.With(l).Add(value)
}

