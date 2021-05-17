package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", c.Value)
}

func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sId, _ := uuid.NewV4()
		s := sId.String()
		s = s + "|" + "sss" + "|" + "fff" + "|" + "rrr"
		fmt.Print(s)
		c = &http.Cookie{
			Name:  "session",
			Value: s,
		}
		http.SetCookie(res, c)
	}
	return c

}
