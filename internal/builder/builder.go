package builder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/azrod/go-prometheus-metrics-builder/types"
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

	IterateStruct(t, v, prefixMetric)

	return s
}

func IterateStruct(t reflect.Type, v reflect.Value, prefixMetric string) {

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
			if field.Tag.Get("name") != "" {
				pm = fmt.Sprintf("%s_%s", prefixMetric, strings.ToLower(field.Tag.Get("name")))
			} else {
				pm = fmt.Sprintf("%s_%s", prefixMetric, strings.ToLower(field.Name))
			}

			IterateStruct(field.Type, v.Field(i), pm)
		case reflect.Ptr:
			InitMetric(field, v.Field(i), prefixMetric)
		}

	}
}

func InitMetric(t reflect.StructField, v reflect.Value, prefixMetric string) {
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
		Labels:    strings.Split(t.Tag.Get("labels"), ","), // split string by comma
	}

	// initialize the metric base
	mb.Init(prefixMetric)

	vv, ok := reflect.New(t.Type.Elem()).Interface().(types.Initializable)
	if !ok {
		panic(fmt.Sprintf("unsupported type %s", t.Type.String()))
	}
	v.Set(reflect.ValueOf(vv.Init(mb)))
}
