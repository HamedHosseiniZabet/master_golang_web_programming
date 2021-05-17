package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string
	Firstname string
	Lastname  string
	Password  []byte
}
type Product struct {
	Id    string
	Name  string
	Color string
}

var tpl *template.Template
var dbSessions = map[string]string{} // SessionID ==> Username
var dbUsers = map[string]User{}      // Username  ==> User
var dbProduct = []Product{}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
	pass, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)
	dbUsers["harpy.eagle.z30@gmail.com"] = User{Username: "harpy.eagle.z30@gmail.com", Firstname: "hamed", Lastname: "hosseini zabet", Password: pass}
	dbProduct = append(dbProduct, Product{getUUID(), "CPU", "GOLD"})
	dbProduct = append(dbProduct, Product{getUUID(), "RAM", "Black"})
	dbProduct = append(dbProduct, Product{getUUID(), "GPU", "Greem"})
	dbProduct = append(dbProduct, Product{getUUID(), "VGA", "Silver"})
	dbProduct = append(dbProduct, Product{getUUID(), "SSD", "Blue"})
	dbProduct = append(dbProduct, Product{getUUID(), "Monitor", "Gray"})
	dbProduct = append(dbProduct, Product{getUUID(), "PC", "White"})
	dbProduct = append(dbProduct, Product{getUUID(), "Laptop", "Red"})
	dbProduct = append(dbProduct, Product{getUUID(), "Mouse", "Black"})
	dbProduct = append(dbProduct, Product{getUUID(), "Headset", "Yellow"})

}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/addProduct", addProduct)
	http.HandleFunc("/getProducts", getProducts)
	http.HandleFunc("/getProduct", getProduct)
	http.HandleFunc("/editProduct", editProduct)
	http.HandleFunc("/deleteProduct", deleteProduct)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

//******************** GET UUID ********************//
func getUUID() string {
	id, _ := uuid.NewV4()
	return id.String()
}

//******************** END GET UUID ********************//

//******************** DELETE PRODUCT ********************//
func deleteProduct(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	id := req.FormValue("id")
	var data Product
	for i, v := range dbProduct {
		if v.Id == id {
			dbProduct = RemoveIndex(dbProduct, i)
			data = v
			break
		}
	}
	jsonProduct, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonProduct)

}

//******************** REMOVE FROM SLICE ********************//
func RemoveIndex(s []Product, index int) []Product {
	return append(s[:index], s[index+1:]...)
}

//******************** END REMOVE FROM SLICE ********************//

//******************** END DELETE PRODUCT ********************//

//******************** EDIT PRODUCT ********************//
func editProduct(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	id := req.FormValue("id")
	var data Product
	for i, v := range dbProduct {
		if v.Id == id {
			dbProduct[i].Name = req.FormValue("name")
			dbProduct[i].Color = req.FormValue("color")
			data = v
			break
		}
	}
	jsonProduct, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonProduct)

}

//******************** END EDIT PRODUCT ********************//

//******************** GET PRODUCT ********************//
func getProduct(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	id := req.FormValue("id")
	var data Product
	for _, v := range dbProduct {
		if v.Id == id {
			data = v
			break
		}
	}
	jsonProduct, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonProduct)

}

//******************** END GET PRODUCT ********************//

//******************** GET PRODUCTS ********************//
func getProducts(res http.ResponseWriter, req *http.Request) {
	jsonProducts, err := json.Marshal(dbProduct)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonProducts)
}

//******************** END GET PRODUCTS ********************//

//******************** ADD PRODUCT ********************//
func addProduct(res http.ResponseWriter, req *http.Request) {
	var p Product
	n := req.FormValue("name")
	c := req.FormValue("color")
	id, _ := uuid.NewV4()
	p = Product{id.String(), n, c}
	dbProduct = append(dbProduct, p)
	// fmt.Printf("%#v\n", dbProduct)
	fmt.Fprint(res, p)
}

//******************** END ADD PRODUCT ********************//

//******************** LOGIN ********************//
func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/home", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		u, ok := dbUsers[un]
		if !ok {
			tpl.ExecuteTemplate(res, "login.gohtml", "Incorrect Username Or Password")
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			tpl.ExecuteTemplate(res, "login.gohtml", "Incorrect Username Or Password")
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un
		http.Redirect(res, req, "/home", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "login.gohtml", nil) // must direct some where
}

//******************** END LOGIN ********************//

//******************** SINGUP ********************//
func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/home", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		for k := range dbUsers {
			if k == req.FormValue("usernamesignup") {
				tpl.ExecuteTemplate(res, "signup.gohtml", "The Username has already been taken")
				return
			}
		}
	}
	if req.Method == http.MethodPost {
		sid, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    sid.String(),
			HttpOnly: true,
			Secure:   true,
		}
		http.SetCookie(res, c)
		un := req.FormValue("usernamesignup")
		fn := req.FormValue("firstnamesignup")
		ln := req.FormValue("lastnamesignup")
		pw := req.FormValue("passwordsignup")
		dbSessions[sid.String()] = un
		bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		}
		dbUsers[un] = User{un, fn, ln, bs}
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil) // must direct some where
}

//******************** END SINGUP ********************//

//******************** HOME ********************//
func home(res http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/signup", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/signup", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "home.gohtml", u)
}

//******************** END HOME ********************//

//******************** LOGOUT ********************//
func logout(res http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)

	tpl.ExecuteTemplate(res, "login.gohtml", nil) // must direct some where
}

//******************** END LOGOUT ********************//
