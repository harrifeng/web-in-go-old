package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title       string
	Description string
}

const tmpl = `Note - Title: {{.Title}}, Description: {{.Description}}`

func main() {

	note := Note{"text/templates", "Template generates textual output"}

	// Create a new tmplate with a name
	t := template.New("note")

	// Parse some content and generate a template
	t, err := t.Parse(tmpl)

	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	//Aplies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Eecute: ", err)
		return
	}

	fmt.Println()
	os.Exit(0)
}
