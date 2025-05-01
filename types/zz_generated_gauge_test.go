package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestGauge_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_gauge",
		Help: "This is a test gauge",
	}
	gauge := &Gauge{}
	result := gauge.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, gauge.GetName())
	assert.Equal(t, mb.Help, gauge.GetHelp())
	assert.Equal(t, mb.Namespace, gauge.GetNamespace())
	assert.Equal(t, mb.Subsystem, gauge.GetSubsystem())
	assert.Equal(t, mb.Labels, gauge.GetLabels())
	assert.IsType(t, &Gauge{}, result)
}

func TestGauge_Collector(t *testing.T) {
	gauge := &Gauge{}
	collector := gauge.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &gaugeCollector{}, collector)
}

func TestGauge_GetType(t *testing.T) {
	gauge := &Gauge{}
	assert.Equal(t, TypeGauge, gauge.GetType())
}

func TestGauge_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	gauge := &Gauge{}
	mb := &Metric{
		Name: "test_gauge",
		Help: "This is a test gauge",
	}
	gauge.Init(mb)

	collector := gauge.Collector().(*gaugeCollector)
	collector.SetValue(1)

	metricFamilies, _ := registry.PRegistry.Gather()
    
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Gauge.Value)
    
}
