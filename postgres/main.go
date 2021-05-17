package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

var db *sql.DB

func init()  {
	var err error
	db,err=sql.Open("postgres","postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err!=nil{
		panic(err)
	}
	if err=db.Ping();err!=nil{
		panic(err)
	}
	fmt.Println("You Connected to postgres")

}

type Book struct {
	Isbn string   `json:"isbn"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price float64 `json:"price"`
}

func main()  {
	http.HandleFunc("/books",booksIndex)
	http.HandleFunc("/books/show",booksShow)
	http.HandleFunc("books/add",addBook)
	http.ListenAndServe(":8080",nil)

}
func addBook(w http.ResponseWriter,r *http.Request)  {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 64)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = f64

	// insert values
	res, err := db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4);", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	fmt.Println(res)
}
func booksShow(w http.ResponseWriter,r *http.Request)  {
	if r.Method!="GET"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	isbn:=r.FormValue("isbn")
	if isbn==""{
		http.Error(w,http.StatusText(400),http.StatusBadRequest)
		return
	}
	row:=db.QueryRow("SELECT * FROM books WHERE isbn=$1",isbn)
	bk:=Book{}
	err:=row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err==sql.ErrNoRows:
		http.NotFound(w,r)
		return
	case err!=nil:
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	bkJ,err := json.Marshal(bk)
	if err!=nil{
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bkJ)

}

func booksIndex(w http.ResponseWriter,r *http.Request)  {
	if r.Method!="GET"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	rows,err:=db.Query("SELECT * FROM books")
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	bks:=make([]Book,0)
	for rows.Next(){
		bk:=Book{}
		err=rows.Scan(&bk.Isbn,&bk.Title,&bk.Author,&bk.Price)
		if err!=nil{
			http.Error(w,http.StatusText(500),500)
			return
		}
		bks=append(bks,bk)
	}
	if err=rows.Err();err!=nil{
		http.Error(w,http.StatusText(500),500)
		return
	}
	bksJ,err := json.Marshal(bks)
	//for _,bk:=range bks{
	//	fmt.Fprintf(w,"%s , %s , %s , $%.2f\n",bk.Isbn,bk.Title,bk.Author,bk.Price)
	//}
	if err!=nil{
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bksJ)

}
