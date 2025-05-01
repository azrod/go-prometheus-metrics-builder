package types

type (
	MetricType string
)

func (mt MetricType) String() string {
	return string(mt)
}

const (
	// TypeCounter is the type of the metric counter
	TypeCounter MetricType = "counter"
	// TypeGauge is the type of the metric gauge
	TypeGauge MetricType = "gauge"
	// TypeHistogram is the type of the metric histogram
	TypeHistogram MetricType = "histogram"
	// TypeSummary is the type of the metric summary
	TypeSummary MetricType = "summary"

	// TypeCounterVec is the type of the metric counter vector
	TypeCounterVec MetricType = "counterVec"
	// TypeGaugeVec is the type of the metric gauge vector
	TypeGaugeVec MetricType = "gaugeVec"
	// TypeHistogramVec is the type of the metric histogram vector
	TypeHistogramVec MetricType = "histogramVec"
	// TypeSummaryVec is the type of the metric summary vector
	TypeSummaryVec MetricType = "summaryVec"
)

var (
	// MetricTypes is a list of all the metric types
	MetricTypes = []MetricType{
		TypeCounter,
		TypeGauge,
		TypeHistogram,
		TypeSummary,
		TypeCounterVec,
		TypeGaugeVec,
		TypeHistogramVec,
		TypeSummaryVec,
	}

	// MetricTypesVec is a list of all the metric vector types
	MetricTypesVec = []MetricType{
		TypeCounterVec,
		TypeGaugeVec,
		TypeHistogramVec,
		TypeSummaryVec,
	}

	MetricTypesVecStr = func() []string {
		var m []string
		for _, mt := range MetricTypesVec {
			m = append(m, string(mt))
		}
		return m
	}()
)
