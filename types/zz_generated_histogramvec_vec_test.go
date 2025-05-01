package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestHistogramVec_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_histogramVec",
		Help: "This is a test histogramVec",
	}
	histogramVec := &HistogramVec{}
	result := histogramVec.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, histogramVec.GetName())
	assert.Equal(t, mb.Help, histogramVec.GetHelp())
	assert.Equal(t, mb.Namespace, histogramVec.GetNamespace())
	assert.Equal(t, mb.Subsystem, histogramVec.GetSubsystem())
	assert.Equal(t, mb.Labels, histogramVec.GetLabels())
	assert.IsType(t, &HistogramVec{}, result)
}

func TestHistogramVec_GetType(t *testing.T) {
	cv := &HistogramVec{}
	assert.Equal(t, TypeHistogramVec, cv.GetType(), "Expected metric type to be TypeHistogramVec")
}

func TestHistogramVecCollector_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	histogramVec := &HistogramVec{}
	mb := &Metric{
		Name: "test_histogramVec",
		Help: "This is a test histogramVec",
        Labels: []string{"label1", "label2"},
	}
	histogramVec.Init(mb)

	labels := map[string]string{
		"label1": "value1",
		"label2": "value2",
	}

	collector := histogramVec.Collector().(*histogramVecCollector)
	collector.SetValue(1, labels)

	metricFamilies, _ := registry.PRegistry.Gather()
    
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Histogram.SampleSum)
	
}

func TestHistogramVec_Collector(t *testing.T) {
    histogramVec := &HistogramVec{}
	collector := histogramVec.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &histogramVecCollector{}, collector)
}
