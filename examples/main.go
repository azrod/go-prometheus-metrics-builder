package main

import (
	"context"
	"time"

	pmbuilder "github.com/azrod/go-prometheus-metrics-builder"
)

func main() {
	metrics := &demo{
		InstanceInterface: &pmbuilder.DefaultInstance{
			PrefixMetric: "myapp",
		},
	}

	pmbuilder.New(metrics)

	go func() {
		for {
			metrics.API.Redis.Get.Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			metrics.API.DB.Get.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	go metrics.ListenAndServe(context.Background(), ":8080")

	select {}
}
