package main

import (
	"log"
	"net/http"
	"text/template"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(res, "index.gohtml", req.Form)
	if err != nil {
		log.Fatal(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
