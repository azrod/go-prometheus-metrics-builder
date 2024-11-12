package main

import (
	"testing"

	pmbuilder "github.com/azrod/go-prometheus-metrics-builder"
	"github.com/azrod/go-prometheus-metrics-builder/pkg/tests"
	"github.com/azrod/go-prometheus-metrics-builder/types"
)

func Test_newMetric(t *testing.T) {
	d := &demo{
		InstanceInterface: &pmbuilder.DefaultInstance{
			PrefixMetric: "unittest",
		},
	}

	pmbuilder.New(d)

	metrics := tests.Helper(d)

	// * Tests all the metrics in the struct
	for _, tm := range types.MetricTypes {
		if v, ok := metrics[tm]; ok {
			for _, metric := range v {
				metric.Run(t)
			}
		}
	}

}
