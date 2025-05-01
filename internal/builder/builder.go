package builder

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/azrod/go-prometheus-metrics-builder/types"
)

const (
	// ignoreTag is used to ignore the field
	ignoreTag = "_"
)

func BuildMetrics[T any](s T, prefixMetric string, racine bool) T {
	var (
		t reflect.Type
		v reflect.Value
	)

	// Check if the main struct is a pointer
	if racine && reflect.TypeOf(s).Kind() != reflect.Ptr {
		panic("Register only accepts pointers")
	}

	switch reflect.TypeOf(s).Kind() {
	case reflect.Ptr:
		t = reflect.TypeOf(s).Elem()
		v = reflect.ValueOf(s).Elem()
	case reflect.Struct:
		t = reflect.TypeOf(s)
		v = reflect.ValueOf(&s).Elem()

		// Check if the struct is initialized
		if v.IsNil() {
			// Initialize the struct
			v.Set(reflect.New(t))
		}
	}

	IterateStruct(t, v, prefixMetric, []string{})

	return s
}

func IterateStruct(t reflect.Type, v reflect.Value, prefixMetric string, labels []string) {

	// Iterate over the struct fields
	for i := 0; i < t.NumField(); i++ {
		// Get the field
		field := t.Field(i)

		switch field.Type.Kind() {
		case reflect.Struct:
			if field.Type.String() == "pmbuilder.InstanceInterface" {
				continue
			}

			// build the new prefixMetric
			// example:
			// prefixMetric = "myApp"
			// type metric struct {
			// 	API struct {
			// 		DB struct {
			// 			Get *types.Counter `help:"Demo counter"`
			// 			Set *types.Counter `help:"Demo counter"`
			// 		}
			//  }
			// }
			// new prefixMetric = "myApp_api_db" for API.DB.Get and API.DB.Set
			// special case if the struct have a tag name defined in the struct field
			// example:
			// type metric struct {
			// 	API struct {
			// 		DB struct {
			// 			Get *types.Counter `help:"Demo counter" name:"get"`
			// 			Set *types.Counter `help:"Demo counter" name:"set"`
			// 		} `name:"database"`
			//  }
			// }
			// new prefixMetric = "myApp_api_database" for API.DB.Get and API.DB.Set
			var pm string

			switch field.Tag.Get("name") {
			case "":
				pm = fmt.Sprintf("%s_%s", prefixMetric, strings.ToLower(field.Name))
			case ignoreTag:
				pm = prefixMetric
			default:
				pm = fmt.Sprintf("%s_%s", prefixMetric, strings.ToLower(field.Tag.Get("name")))
			}

			// build the common labels
			// example:
			// type metric struct {
			// 	API struct {
			// 		DB struct {
			// 			Get *types.CounterVec `help:"Demo counter" labels:"server,version"`
			// 			Set *types.CounterVec `help:"Demo counter" labels:"server,version"`
			// 		}
			//  }
			// }
			// labels = []string{"server", "version"}
			//
			// special case if the struct have a tag labels defined in the struct field
			// all labels are added to each metric
			// example:
			// type metric struct {
			// 	API struct {
			// 		DB struct {
			// 			Get *types.CounterVec `help:"Demo counter"`
			// 			Set *types.CounterVec `help:"Demo counter"`
			// 		} `name:"database" labels:"server,version"`
			//  }
			//  }
			// }
			// labels = []string{"server", "version"}
			var commonLabels []string
			if field.Tag.Get("labels") != "" {
				commonLabels = strings.Split(field.Tag.Get("labels"), ",")
			}
			// merge the common labels with the labels from the parent struct
			labels = append(labels, commonLabels...)

			IterateStruct(field.Type, v.Field(i), pm, labels)
		case reflect.Ptr:
			InitMetric(field, v.Field(i), prefixMetric, labels)
		}

	}
}

func InitMetric(t reflect.StructField, v reflect.Value, prefixMetric string, labels []string) {
	mb := &types.Metric{
		Name: func() string {
			// If the name is not provided, use the field name
			if t.Tag.Get("name") == "" {
				return t.Name
			}

			return t.Tag.Get("name")
		}(),
		Help: func() string {
			if t.Tag.Get("help") == "" {
				panic(fmt.Sprintf("help is required for %s", t.Name))
			}

			return t.Tag.Get("help")
		}(),
		Namespace: t.Tag.Get("namespace"),
		Subsystem: t.Tag.Get("subsystem"),
		Labels: func() (l []string) {
			l = labels
			if t.Tag.Get("labels") != "" {
				l = append(l, strings.Split(t.Tag.Get("labels"), ",")...) // split string by comma
			}
			return l
		}(),
	}

	// initialize the metric base
	mb.Init(prefixMetric)

	vv, ok := reflect.New(t.Type.Elem()).Interface().(types.Initializable)
	if !ok {
		panic(fmt.Sprintf("unsupported type %s", t.Type.String()))
	}

	// If the metric is a Vec, we need to check if the labels are provided
	// If the metric is a Vec and the labels are not provided, we need to panic
	if slices.Contains(types.MetricTypesVecStr, vv.GetType().String()) && len(mb.Labels) == 0 {
		panic(fmt.Sprintf("labels are required for %s", mb.Name))
	}

	v.Set(reflect.ValueOf(vv.Init(mb)))
}
