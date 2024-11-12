package types

import "github.com/prometheus/client_golang/prometheus"

type (
	Collector interface {
		MetricInterface
		prometheus.Collector

		// SetValue sets the value of the metric.
		SetValue(value float64, labels ...map[string]string)
	}
)
