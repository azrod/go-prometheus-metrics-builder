package main

import (
	pmbuilder "github.com/azrod/go-prometheus-metrics-builder"
	"github.com/azrod/go-prometheus-metrics-builder/types"
)

type demo struct {
	pmbuilder.InstanceInterface
	API struct {
		DB struct {
			Get *types.Counter `help:"Demo counter"`
			Set *types.Counter `help:"Demo counter"`
		} `name:"database"`
		Redis struct {
			Get *types.Counter `help:"Demo counter"`
			Set *types.Counter `help:"Demo counter"`
		}
	}
}
