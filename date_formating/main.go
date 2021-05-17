package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format(time.ANSIC)
}

var fm = template.FuncMap{
	"df": monthDayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
