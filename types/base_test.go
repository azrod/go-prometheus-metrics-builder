package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetric_Init(t *testing.T) {
	tests := []struct {
		name         string
		metric       Metric
		prefixMetric string
		expectedName string
	}{
		{
			name: "simple metric",
			metric: Metric{
				Name: "test_metric",
			},
			prefixMetric: "prefix",
			expectedName: "prefix_test_metric",
		},
		{
			name: "metric with namespace and subsystem",
			metric: Metric{
				Name:      "test_metric",
				Namespace: "test_namespace",
				Subsystem: "test_subsystem",
			},
			prefixMetric: "prefix",
			expectedName: "prefix_test_metric",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.metric.Init(tt.prefixMetric)
			assert.Equal(t, tt.expectedName, tt.metric.Name)
		})
	}
}
