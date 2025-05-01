package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestHistogram_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_histogram",
		Help: "This is a test histogram",
	}
	histogram := &Histogram{}
	result := histogram.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, histogram.GetName())
	assert.Equal(t, mb.Help, histogram.GetHelp())
	assert.Equal(t, mb.Namespace, histogram.GetNamespace())
	assert.Equal(t, mb.Subsystem, histogram.GetSubsystem())
	assert.Equal(t, mb.Labels, histogram.GetLabels())
	assert.IsType(t, &Histogram{}, result)
}

func TestHistogram_Collector(t *testing.T) {
	histogram := &Histogram{}
	collector := histogram.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &histogramCollector{}, collector)
}

func TestHistogram_GetType(t *testing.T) {
	histogram := &Histogram{}
	assert.Equal(t, TypeHistogram, histogram.GetType())
}

func TestHistogram_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	histogram := &Histogram{}
	mb := &Metric{
		Name: "test_histogram",
		Help: "This is a test histogram",
	}
	histogram.Init(mb)

	collector := histogram.Collector().(*histogramCollector)
	collector.SetValue(1)

	metricFamilies, _ := registry.PRegistry.Gather()
    
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Histogram.SampleSum)
	
}
