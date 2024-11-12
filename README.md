<div align="center">
    <a href="https://github.com/azrod/go-prometheus-metrics-builder/releases/latest">
      <img alt="Latest release" src="https://img.shields.io/github/v/release/azrod/go-prometheus-metrics-builder?style=for-the-badge&logo=starship&color=C9CBFF&logoColor=D9E0EE&labelColor=302D41&include_prerelease&sort=semver" />
    </a>
    <a href="https://github.com/azrod/go-prometheus-metrics-builder/pulse">
      <img alt="Last commit" src="https://img.shields.io/github/last-commit/azrod/go-prometheus-metrics-builder?style=for-the-badge&logo=starship&color=8bd5ca&logoColor=D9E0EE&labelColor=302D41"/>
    </a>
    <a href="https://github.com/azrod/go-prometheus-metrics-builder/blob/main/LICENSE">
      <img alt="License" src="https://img.shields.io/github/license/azrod/go-prometheus-metrics-builder?style=for-the-badge&logo=starship&color=ee999f&logoColor=D9E0EE&labelColor=302D41" />
    </a>
    <a href="https://github.com/azrod/go-prometheus-metrics-builder/stargazers">
      <img alt="Stars" src="https://img.shields.io/github/stars/azrod/go-prometheus-metrics-builder?style=for-the-badge&logo=starship&color=c69ff5&logoColor=D9E0EE&labelColor=302D41" />
    </a>
    <a href="https://github.com/azrod/go-prometheus-metrics-builder/issues">
      <img alt="Issues" src="https://img.shields.io/github/issues/azrod/go-prometheus-metrics-builder?style=for-the-badge&logo=bilibili&color=F5E0DC&logoColor=D9E0EE&labelColor=302D41" />
    </a>
</div>

# Go Prometheus Metrics Builder

> [!CAUTION]
> This project is in early development and is not yet ready for production use. You are welcome to try it out and provide feedback, but be aware that the library may change at any time.

This library provides a simple way to create Prometheus metrics in Go. It is designed to be easy to use and to provide a simple way to create metrics without having to worry about the underlying details of the Prometheus client library.

## Example

```go
package main

import (
  "context"

  pmbuilder "github.com/azrod/go-prometheus-metrics-builder"
)

type demo struct {
  pmbuilder.InstanceInterface
  API struct {
    DB struct {
      Get *types.Counter `help:"Database get counter"`
      Set *types.Counter `help:"Database set counter"`
    } `name:"database"`
    Redis struct {
      Get *types.Counter `help:"Redis get counter"`
      Set *types.Counter `help:"Redis set counter"`
    }
  }
}
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

  go metrics.ListenAndServe(context.Background(), ":8080")

  select {}
}
```

In this example, we create a new instance of the `demo` struct, which implements the `InstanceInterface` interface. We then create two counters, `Get` and `Set`, for the `database` and `Redis` APIs. We then create a new `demo` instance and start incrementing the `Get` counter for the `Redis` API every second. Finally, we start the metrics server on port 8080.

```shell
$ curl localhost:8080/metrics
[...]
# HELP myapp_api_database_get Database get counter
# TYPE myapp_api_database_get counter
myapp_api_database_get 2
# HELP myapp_api_database_set Database set counter
# TYPE myapp_api_database_set counter
myapp_api_database_set 0
# HELP myapp_api_redis_get Redis get counter
# TYPE myapp_api_redis_get counter
myapp_api_redis_get 0
# HELP myapp_api_redis_set Redis set counter
# TYPE myapp_api_redis_set counter
myapp_api_redis_set 0
[...]
```

## Installation

To install the library, you can use `go get`:

```shell
go get github.com/azrod/go-prometheus-metrics-builder
```

## Usage

To use the library, you need to create a struct that implements the `InstanceInterface` interface. pmbuilder provides a default implementation of this interface, `DefaultInstance`, which you can embed in your struct to get the default behavior.

```go
type demo struct {
  pmbuilder.InstanceInterface
}
```

You can then define the metrics you want to create as fields in your struct. The type of the field should be a pointer to one of the metric types provided by the library, such as `Counter`, `Gauge`, `Summary` or `Histogram`.

```go
type demo struct {
  pmbuilder.InstanceInterface
  API struct {
    DB struct {
      Get *types.Counter `help:"Database get counter"`
      Set *types.Counter `help:"Database set counter"`
    } `name:"database"`
    Redis struct {
      Get *types.Counter `help:"Redis get counter"`
      Set *types.Counter `help:"Redis set counter"`
    }
  }
}
```

You can then create a new instance of your struct and call the `New` function to create the metrics.

```go
metrics := &demo{
  InstanceInterface: &pmbuilder.DefaultInstance{
    PrefixMetric: "myapp",
  },
}

pmbuilder.New(metrics)
```

### Tags in the struct

A following golang tags are available to customize the metrics:

- `name`: The name of the metric. If not provided, the name of the field will be used.
- `help`: The help text for the metric. **Required.**
- `labels`: A comma-separated list of labels for the metric. The labels will be added to the metric as tags.
- `namespace`: The namespace for the metric. Used to group metrics together.
- `subsystem`: The subsystem for the metric. Used to further categorize metrics within a namespace.

All tags are optional, except for `help`.

A special tag `name` is available to customize the name of the metric. If not provided, the name of the field will be used. This is useful when you want to use a different name for the metric than the name of the field.

```go
type demo struct {
  pmbuilder.InstanceInterface
  API struct {
    DB struct {
      Get *types.Counter `help:"Database get counter" name:"get_counter"`
      Set *types.Counter `help:"Database set counter" name:"set_counter"`
    } `name:"database"`
    [...]
  }
}
```

In this example, we use the `name` tag on the `DB` struct to set the name of the metric to `database`. We also use the `name` tag on the `Get` and `Set` fields to set the name of the metrics to `get_counter` and `set_counter`, respectively.

Name generated : `myapp_api_database_get_counter` and `myapp_api_database_set_counter`

### Labels

Not yet implemented.

## License

This project is licensed under the Apache2 License - see the [LICENSE](LICENSE) file for details.
