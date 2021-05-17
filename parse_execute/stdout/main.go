package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	var tpl *template.Template
	var err error
	tpl, err = template.ParseFiles("templates/tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

}
