package main

import (
	"flag"
	"io/ioutil"
	"os"
	"text/template"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Data interface{}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	templateFilePathPtr := flag.String("t", "example/custom.gotmpl", "(Required) Source template file.")
	dataFilePathPtr := flag.String("d", "example/data.yaml", "(Required) Data in yaml format.")
	outFilePathPtr := flag.String("o", "", "(Optional) Data in yaml format.")

	flag.Parse()

	var data Data

	dataTxt, err := ioutil.ReadFile(*dataFilePathPtr)
	check(err)

	// Load DataTxt as Yaml Data object
	err = yaml.Unmarshal(dataTxt, &data)
	check(err)

	// Get template format
	tmpl, err := ioutil.ReadFile(*templateFilePathPtr)
	check(err)

	// Setup Templating
	t := template.New("Template").Funcs(TemplateFuncMap)
	t, err = t.Parse(string(tmpl))
	check(err)

	if *outFilePathPtr == "" {
		// Execute Template and print
		err = t.Execute(os.Stdout, data)
		check(err)
	} else {
		// When outFile exist, write to the file path
		f, err := os.Create(*outFilePathPtr)
		check(err)
		err = t.Execute(f, data)
		check(err)
		f.Close()
	}
}
