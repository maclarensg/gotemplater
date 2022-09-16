package main

import (
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"gopkg.in/yaml.v2"
)

func TestTemplateFunc(t *testing.T) {
	var data Data
	dataTxt, err := ioutil.ReadFile("example/data.yaml")
	check(err)

	// Load DataTxt as Yaml Data object
	err = yaml.Unmarshal(dataTxt, &data)
	check(err)

	// Get template format
	tmpl, err := ioutil.ReadFile("example/custom.txt.gotmpl")
	check(err)

	// Setup Templating
	template := template.New("Template").Funcs(TemplateFuncMap)
	template, err = template.Parse(string(tmpl))
	check(err)
	err = template.Execute(os.Stdout, data)
	check(err)
}
