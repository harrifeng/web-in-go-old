package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {

	t, err := template.New("test").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('XSS Injection')</script>")
	if err != nil {
		log.Fatal("Execute: ", err)
	}
	fmt.Println()
	os.Exit(0)
}
