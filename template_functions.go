package gogi

import (
	"bytes"
	"html/template"
	"log"
)

var templateFunctions = template.FuncMap{
	"exec": func(t *template.Template, d interface{}) template.HTML {
		var tpl bytes.Buffer
		err := t.Execute(&tpl, d)
		if err != nil {
			log.Fatal(err)
		}
		return template.HTML(tpl.String())
	},
}
