package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Note struct {
	Title       string
	Description string
}

const tmpl = `Note are:
{{range . }}
        Title: {{.Title}}, Description: {{.Description}}
{{end}}
`

func main() {

	notes := []Note{
		{"text/template", "Template generates textual output"},
		{"html/template", "Template generates HTML output"},
	}

	t := template.New("note")
	t, err := t.Parse(tmpl)

	if err != nil {
		log.Fatal("Parse:", err)
		return
	}

	if err := t.Execute(os.Stdout, notes); err != nil {
		log.Fatal("Execute:", err)
		return
	}
	fmt.Println()
	os.Exit(0)
}
