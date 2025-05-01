package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestGaugeVec_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_gaugeVec",
		Help: "This is a test gaugeVec",
	}
	gaugeVec := &GaugeVec{}
	result := gaugeVec.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, gaugeVec.GetName())
	assert.Equal(t, mb.Help, gaugeVec.GetHelp())
	assert.Equal(t, mb.Namespace, gaugeVec.GetNamespace())
	assert.Equal(t, mb.Subsystem, gaugeVec.GetSubsystem())
	assert.Equal(t, mb.Labels, gaugeVec.GetLabels())
	assert.IsType(t, &GaugeVec{}, result)
}

func TestGaugeVec_GetType(t *testing.T) {
	cv := &GaugeVec{}
	assert.Equal(t, TypeGaugeVec, cv.GetType(), "Expected metric type to be TypeGaugeVec")
}

func TestGaugeVecCollector_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	gaugeVec := &GaugeVec{}
	mb := &Metric{
		Name: "test_gaugeVec",
		Help: "This is a test gaugeVec",
        Labels: []string{"label1", "label2"},
	}
	gaugeVec.Init(mb)

	labels := map[string]string{
		"label1": "value1",
		"label2": "value2",
	}

	collector := gaugeVec.Collector().(*gaugeVecCollector)
	collector.SetValue(1, labels)

	metricFamilies, _ := registry.PRegistry.Gather()
    
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Gauge.Value)
    
}

func TestGaugeVec_Collector(t *testing.T) {
    gaugeVec := &GaugeVec{}
	collector := gaugeVec.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &gaugeVecCollector{}, collector)
}
