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

var funcMap = template.FuncMap{
	"toLower": func(s string) string {
		return strcase.ToLowerCamel(s)
	},
}

type templateData struct {
	MetricType string
}

func main() {

	tmpl, err := template.New("type.tmpl").Funcs(funcMap).ParseFiles("tools/types_template/type.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range listOfTypes {

		// write template to file
		f, err := os.Create(fmt.Sprintf("types/zz_generated_%s.go", strings.ToLower(t)))
		if err != nil {
			log.Default().Printf("Failed to create file: %v", err)
			f.Close()
			os.Exit(1) //nolint: gocritic
		}
		defer f.Close()

		tD := templateData{
			MetricType: t,
		}

		if err := tmpl.Execute(f, tD); err != nil {
			log.Default().Printf("Failed to execute template: %v", err)
			os.Exit(1)
		}
	}

}
