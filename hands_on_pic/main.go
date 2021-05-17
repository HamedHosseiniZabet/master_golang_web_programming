package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}
func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}
func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(res, "dog.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
