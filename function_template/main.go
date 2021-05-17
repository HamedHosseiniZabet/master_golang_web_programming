package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))

}

func main() {
	b := sage{
		Name:  "Buddah",
		Motto: "buddah bio",
	}
	g := sage{
		Name:  "gandi",
		Motto: "gandi bio",
	}
	m := sage{
		Name:  "mlk",
		Motto: "mlk bio",
	}
	sages := []sage{b, g, m}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
