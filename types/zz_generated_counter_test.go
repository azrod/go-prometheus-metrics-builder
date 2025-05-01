package types

import (
	"testing"

	"github.com/azrod/go-prometheus-metrics-builder/pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestCounter_Init(t *testing.T) {
	mb := &Metric{
		Name: "test_counter",
		Help: "This is a test counter",
	}
	counter := &Counter{}
	result := counter.Init(mb)

	assert.NotNil(t, result)
	assert.Equal(t, mb.Name, counter.GetName())
	assert.Equal(t, mb.Help, counter.GetHelp())
	assert.Equal(t, mb.Namespace, counter.GetNamespace())
	assert.Equal(t, mb.Subsystem, counter.GetSubsystem())
	assert.Equal(t, mb.Labels, counter.GetLabels())
	assert.IsType(t, &Counter{}, result)
}

func TestCounter_Collector(t *testing.T) {
	counter := &Counter{}
	collector := counter.Collector()

	assert.NotNil(t, collector)
	assert.IsType(t, &counterCollector{}, collector)
}

func TestCounter_GetType(t *testing.T) {
	counter := &Counter{}
	assert.Equal(t, TypeCounter, counter.GetType())
}

func TestCounter_SetValue(t *testing.T) {
	registry.PRegistry = prometheus.NewRegistry()
	counter := &Counter{}
	mb := &Metric{
		Name: "test_counter",
		Help: "This is a test counter",
	}
	counter.Init(mb)

	collector := counter.Collector().(*counterCollector)
	collector.SetValue(1)

	metricFamilies, _ := registry.PRegistry.Gather()
    
	assert.Equal(t, 1.0, *metricFamilies[0].Metric[0].Counter.Value)
    
}
