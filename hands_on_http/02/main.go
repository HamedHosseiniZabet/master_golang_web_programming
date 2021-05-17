package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(foo))
	mux.Handle("/dog/", http.HandlerFunc(bar))
	mux.Handle("/me/", http.HandlerFunc(myName))
	http.ListenAndServe(":8080", mux)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "Inside foo Route (/)")
	if err != nil {
		log.Fatalln(err)
	}

}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func myName(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello mcleod")
}
