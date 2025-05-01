package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var listOfTypes = []string{
	"Counter",
	"Gauge",
	"Histogram",
	"Summary",
}

var listOfTmpl = []string{
	"type.tmpl",
	"type_test.tmpl",
	"type_vec.tmpl",
	"type_vec_test.tmpl",
}

var funcMap = template.FuncMap{
	"toLower": func(s string) string {
		return strcase.ToLowerCamel(s)
	},
}

type templateData struct {
	MetricType       string
	MetricWithoutVec string
	IsVec            bool
}

func main() {

	for _, tmplFile := range listOfTmpl {

		tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(fmt.Sprintf("tools/types_template/%s", tmplFile))
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, t := range listOfTypes {

			var (
				f   *os.File
				err error
			)

			tD := templateData{}

			// write template to file
			switch {
			case tmplFile == "type.tmpl":
				f, err = os.Create(fmt.Sprintf("types/zz_generated_%s.go", strings.ToLower(t)))
			case tmplFile == "type_test.tmpl":
				f, err = os.Create(fmt.Sprintf("types/zz_generated_%s_test.go", strings.ToLower(t)))
			case tmplFile == "type_vec.tmpl":
				t = fmt.Sprintf("%sVec", t)
				f, err = os.Create(fmt.Sprintf("types/zz_generated_%s_vec.go", strings.ToLower(t)))
				tD.IsVec = true
				tD.MetricWithoutVec = t[:len(t)-3]
			case tmplFile == "type_vec_test.tmpl":
				t = fmt.Sprintf("%sVec", t)
				f, err = os.Create(fmt.Sprintf("types/zz_generated_%s_vec_test.go", strings.ToLower(t)))
				tD.IsVec = true
				tD.MetricWithoutVec = t[:len(t)-3]
			}
			if err != nil {
				log.Default().Printf("Failed to create file: %v", err)
				f.Close()
				os.Exit(1) //nolint: gocritic
			}
			defer f.Close()

			tD.MetricType = t
			if tD.MetricWithoutVec == "" {
				tD.MetricWithoutVec = t
			}

			if err := tmpl.Execute(f, tD); err != nil {
				log.Default().Printf("Failed to execute template: %v", err)
				os.Exit(1)
			}
		}
	}

}
