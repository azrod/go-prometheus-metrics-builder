package types

type (
	MetricType string
)

const (
	// TypeCounter is the type of the metric counter
	TypeCounter MetricType = "counter"
	// TypeGauge is the type of the metric gauge
	TypeGauge MetricType = "gauge"
	// TypeHistogram is the type of the metric histogram
	TypeHistogram MetricType = "histogram"
	// TypeSummary is the type of the metric summary
	TypeSummary MetricType = "summary"
)

var (
	// MetricTypes is a list of all the metric types
	MetricTypes = []MetricType{
		TypeCounter,
		TypeGauge,
		TypeHistogram,
		TypeSummary,
	}
)
