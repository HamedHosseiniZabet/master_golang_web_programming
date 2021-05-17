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
		Motto: "buddah bio",
	}
	gandi := sage{
		Name:  "gandi",
		Motto: "gandi bio",
	}
	mlk := sage{
		Name:  "mlk",
		Motto: "mlk bio",
	}
	Mohamaad := sage{
		Name:  "Mohamaad",
		Motto: "Mohamaad bio",
	}

	sages := []sage{Buddah, gandi, mlk, Mohamaad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
