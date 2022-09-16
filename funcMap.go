package main

import (
	"strings"
	"text/template"
)

var TemplateFuncMap = template.FuncMap{
	"inc": func(n int) int {
		return n + 1
	},
	"split": func(d string, s string) []string {
		return strings.Split(s, d)
	},
}
