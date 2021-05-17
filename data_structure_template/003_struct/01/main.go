package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	Buddah := sage{
		Name:  "Buddah",
		Motto: "some text",
	}

	err := tpl.Execute(os.Stdout, Buddah)
	if err != nil {
		log.Fatalln(err)
	}
}
