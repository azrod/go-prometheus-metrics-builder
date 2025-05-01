package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestCounterVec_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_counterVec",
		Help: "This is a test counterVec",
	}
	counterVec := &CounterVec{}
	result := counterVec.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, counterVec.GetName())
	assert.Equal(t, mb.Help, counterVec.GetHelp())
	assert.Equal(t, mb.Namespace, counterVec.GetNamespace())
	assert.Equal(t, mb.Subsystem, counterVec.GetSubsystem())
	assert.Equal(t, mb.Labels, counterVec.GetLabels())
	assert.IsType(t, &CounterVec{}, result)
}

func TestCounterVec_GetType(t *testing.T) {
	cv := &CounterVec{}
	assert.Equal(t, TypeCounterVec, cv.GetType(), "Expected metric type to be TypeCounterVec")
}

func TestCounterVecCollector_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	counterVec := &CounterVec{}
	mb := &Metric{
		Name: "test_counterVec",
		Help: "This is a test counterVec",
        Labels: []string{"label1", "label2"},
	}
	counterVec.Init(mb)

	labels := map[string]string{
		"label1": "value1",
		"label2": "value2",
	}

	collector := counterVec.Collector().(*counterVecCollector)
	collector.SetValue(1, labels)

	metricFamilies, _ := registry.PRegistry.Gather()
     
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Counter.Value)
    
}

func TestCounterVec_Collector(t *testing.T) {
    counterVec := &CounterVec{}
	collector := counterVec.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &counterVecCollector{}, collector)
}
