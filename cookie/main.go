package main

import (
	"fmt"
	"log"
	"net/http"
)

var num int = 1

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: fmt.Sprint(num),
	})
	fmt.Fprintln(res, "Cookie written - check your browser")
	num++
}
func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Fprintln(res, "your cookie: ", c)
	}
}
