package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/master_go_programming/crud_v2/config"
	"github.com/master_go_programming/crud_v2/helpers"
	"github.com/master_go_programming/crud_v2/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type HomeController struct {
	session *mgo.Session
}

func NewHomeController(s *mgo.Session) *HomeController {
	return &HomeController{s}
}

func (hc HomeController) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	v,_:=helpers.AlreadyLoggedIn(r)
	if v==true{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", nil) // must direct some where
}
func (hc HomeController) SignUpProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	u.Username = r.FormValue("username")
	u.Firstname = r.FormValue("firstname")
	u.Lastname = r.FormValue("lastname")
	u.Password = r.FormValue("password")
	u.Id = bson.NewObjectId()
	err:=hc.session.DB("demo").C("users").Insert(u)
	if err!=nil{
		panic(err)
	}
	sid := bson.NewObjectId()
	uid := u.Id
	s := models.Session{Id: sid, UserId: uid}
	err=hc.session.DB("demo").C("sessions").Insert(s)
	if err!=nil{
		panic(err)
	}
	c := &http.Cookie{
		Name:  "session",
		Value: sid.Hex(), // object id to string
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, c)
	http.Redirect(w,r,"/",http.StatusSeeOther)
}

func (hc HomeController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	v,_:=helpers.AlreadyLoggedIn(r)
	if v==true{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "login.gohtml", nil)
}
func (hc HomeController) LoginProcess(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	u:=models.User{}
	un:=r.FormValue("username")
	p:=r.FormValue("password")
	err:=hc.session.DB("demo").C("users").Find(bson.M{"username":un}).One(&u)
	if err!=nil{
		config.TPL.ExecuteTemplate(w,"login.gohtml","Incorrect Username or Password")
		return
	}
	if p!=u.Password{
		config.TPL.ExecuteTemplate(w,"login.gohtml","Incorrect Username or Password")
		return
	}
	s:=models.Session{}
	s.Id=bson.NewObjectId()
	s.UserId=u.Id
	err=hc.session.DB("demo").C("sessions").Insert(s)
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	c:=&http.Cookie{
		Name: "session",
		Value: s.Id.Hex(),
		Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(w,c)
	http.Redirect(w,r,"/",http.StatusSeeOther)
}
func (hc HomeController) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c,err:=r.Cookie("session")
	if err!=nil{
		http.Error(w,http.StatusText(500),500)
		return
	}
	err=hc.session.DB("demo").C("sessions").Remove(bson.M{"_id":bson.ObjectIdHex(c.Value)})
	if err!=nil{
		fmt.Println("hahahaha")
		http.Error(w,http.StatusText(500),500)
		return
	}
	c=&http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w,c)
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	v,_:=helpers.AlreadyLoggedIn(r)
	if v==false{
		http.Redirect(w,r,"/signup",http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "home.gohtml", nil)
}
