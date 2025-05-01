package types

import (
	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	SummaryVec struct {
		Metric
		*prometheus.SummaryVec
	}

	summaryVecCollector struct {
		SummaryVec
	}
)

var _ Initializable = &SummaryVec{}

// Init initializes (or re-initializes) the SummaryVec metric
// The metrics is automatically initialized by the Register function
func (c *SummaryVec) Init(mb *Metric) any {
	c.Metric = *mb
    c.SummaryVec = promauto.With(registry.PRegistry).NewSummaryVec(prometheus.SummaryOpts(c.autoBuildSummaryOpts()), mb.Labels)
	return c
}

// Collector returns the prometheus collector
func (c *SummaryVec) Collector() Collector {
	return &summaryVecCollector{
        SummaryVec: *c,
    }
}

// GetType returns the metric type
func (c *SummaryVec) GetType() MetricType {
	return TypeSummaryVec
}

// Run runs the collector.
func (c *summaryVecCollector) SetValue(value float64, labels ...map[string]string) {

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
    c.SummaryVec.With(l).Observe(value)
}

