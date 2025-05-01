package types

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/prometheus/client_golang/prometheus"
)

// Metric represents a Prometheus metric with its associated metadata.
// Name is the name of the metric.
// Help is a brief description of the metric's purpose.
// Namespace is the namespace of the metric, used to group related metrics.
// Subsystem is the subsystem of the metric, used to further categorize metrics within a namespace.
// Labels are the labels associated with the metric, used to add dimensions to the metric.
type Metric struct {
	Name      string
	Help      string
	Namespace string
	Subsystem string
	Labels    []string
}

type (
	Initializable interface {
		MetricInterface
		Init(*Metric) any
		Collector() Collector
		GetType() MetricType
	}

	MetricInterface interface {
		GetName() string
		GetHelp() string
		GetNamespace() string
		GetSubsystem() string
		GetLabels() []string
	}
)

func (mb *Metric) Init(prefixMetric string) {
	mb.Name = fmt.Sprintf(
		"%s_%s",
		strcase.ToSnake(prefixMetric),
		strcase.ToSnake(mb.Name),
	)
}

// autoBuildOpts builds the prometheus Opts for a metric
func (mb Metric) autoBuildOpts() prometheus.Opts {
	return prometheus.Opts{
		Namespace: mb.Namespace,
		Subsystem: mb.Subsystem,
		Name:      mb.Name,
		Help:      mb.Help,
	}
}

// autoBuildSummaryOpts builds the prometheus Opts for a summary metric
func (mb Metric) autoBuildSummaryOpts() prometheus.SummaryOpts {
	return prometheus.SummaryOpts{
		Namespace: mb.Namespace,
		Subsystem: mb.Subsystem,
		Name:      mb.Name,
		Help:      mb.Help,
	}
}

// autoBuildHistogramOpts builds the prometheus Opts for a histogram metric
func (mb Metric) autoBuildHistogramOpts() prometheus.HistogramOpts {
	return prometheus.HistogramOpts{
		Namespace: mb.Namespace,
		Subsystem: mb.Subsystem,
		Name:      mb.Name,
		Help:      mb.Help,
		Buckets:   []float64{0.001, 0.005, 0.01, 0.02, 0.05, 0.1, 0.5, 1, 2, 5, 10},
	}
}

// GetName returns the metric name
func (mb Metric) GetName() string {
	return mb.Name
}

// GetHelp returns the metric help
func (mb Metric) GetHelp() string {
	return mb.Help
}

// GetNamespace returns the metric namespace
func (mb Metric) GetNamespace() string {
	return mb.Namespace
}

// GetSubsystem returns the metric subsystem
func (mb Metric) GetSubsystem() string {
	return mb.Subsystem
}

// GetLabels returns the metric labels
func (mb Metric) GetLabels() []string {
	return mb.Labels
}
