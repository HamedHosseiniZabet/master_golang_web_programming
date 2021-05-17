package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "home route")
}
func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog route")
}
func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hamed Hosseini")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
