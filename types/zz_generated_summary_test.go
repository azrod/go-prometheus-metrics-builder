package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestSummary_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_summary",
		Help: "This is a test summary",
	}
	summary := &Summary{}
	result := summary.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, summary.GetName())
	assert.Equal(t, mb.Help, summary.GetHelp())
	assert.Equal(t, mb.Namespace, summary.GetNamespace())
	assert.Equal(t, mb.Subsystem, summary.GetSubsystem())
	assert.Equal(t, mb.Labels, summary.GetLabels())
	assert.IsType(t, &Summary{}, result)
}

func TestSummary_Collector(t *testing.T) {
	summary := &Summary{}
	collector := summary.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &summaryCollector{}, collector)
}

func TestSummary_GetType(t *testing.T) {
	summary := &Summary{}
	assert.Equal(t, TypeSummary, summary.GetType())
}

func TestSummary_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	summary := &Summary{}
	mb := &Metric{
		Name: "test_summary",
		Help: "This is a test summary",
	}
	summary.Init(mb)

	collector := summary.Collector().(*summaryCollector)
	collector.SetValue(1)

	metricFamilies, _ := registry.PRegistry.Gather()
    
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Summary.SampleSum)
    
}
