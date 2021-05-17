package main

import (
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbSessions = map[string]string{} // sessionID ==> userID
var dbUsers = map[string]user{}      //userID ==> user

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/bar", bar)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./bootstrap-3.4.1-dist"))))
	http.Handle("/favicno.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func login(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
	}
	// if user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	// process from submissions
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstName")
		l := req.FormValue("lastName")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}
	err = tpl.ExecuteTemplate(res, "index.gohtml", u)
	if err != nil {
		log.Fatalln(err)
	}
}
func bar(res http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
