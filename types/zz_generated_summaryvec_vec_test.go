package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestSummaryVec_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_summaryVec",
		Help: "This is a test summaryVec",
	}
	summaryVec := &SummaryVec{}
	result := summaryVec.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, summaryVec.GetName())
	assert.Equal(t, mb.Help, summaryVec.GetHelp())
	assert.Equal(t, mb.Namespace, summaryVec.GetNamespace())
	assert.Equal(t, mb.Subsystem, summaryVec.GetSubsystem())
	assert.Equal(t, mb.Labels, summaryVec.GetLabels())
	assert.IsType(t, &SummaryVec{}, result)
}

func TestSummaryVec_GetType(t *testing.T) {
	cv := &SummaryVec{}
	assert.Equal(t, TypeSummaryVec, cv.GetType(), "Expected metric type to be TypeSummaryVec")
}

func TestSummaryVecCollector_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	summaryVec := &SummaryVec{}
	mb := &Metric{
		Name: "test_summaryVec",
		Help: "This is a test summaryVec",
        Labels: []string{"label1", "label2"},
	}
	summaryVec.Init(mb)

	labels := map[string]string{
		"label1": "value1",
		"label2": "value2",
	}

	collector := summaryVec.Collector().(*summaryVecCollector)
	collector.SetValue(1, labels)

	metricFamilies, _ := registry.PRegistry.Gather()
    
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Summary.SampleSum)
    
}

func TestSummaryVec_Collector(t *testing.T) {
    summaryVec := &SummaryVec{}
	collector := summaryVec.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &summaryVecCollector{}, collector)
}
