package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/master_go_programming/crud_v2/config"
	"github.com/master_go_programming/crud_v2/models"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{session: s}
}

func (uc UserController) AddUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	u.Username = r.FormValue("username")
	u.Firstname = r.FormValue("firstname")
	u.Lastname = r.FormValue("lastname")
	u.Password = r.FormValue("password")
	u.Id = bson.NewObjectId()
	err:=uc.session.DB("demo").C("users").Insert(u)
	if err!=nil{
		panic(err)
	}
	// uj, err := json.Marshal(u)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// }
	// w.Header().Set("Content-Type", "text/html")
	// w.WriteHeader(http.StatusCreated) //201
	//w.Write(uj)
	sid := bson.NewObjectId()
	uid := u.Id
	s := models.Session{Id: sid, UserId: uid}
	err=uc.session.DB("demo").C("sessions").Insert(s)
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
	config.TPL.ExecuteTemplate(w, "home.gohtml", nil)
}


