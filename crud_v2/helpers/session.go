package helpers

import (
	"errors"
	"github.com/master_go_programming/crud_v2/config"
	"github.com/master_go_programming/crud_v2/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetUser(r *http.Request,id string) (models.User,error) {
	u:=models.User{}
	_,err:=r.Cookie("session")
	if err!=nil{
		return u,errors.New("400. Bad Request.")
	}
	oid:=bson.ObjectIdHex(id) // string to objectId
	err=config.GetMongoDbSession().DB("demo").C("users").FindId(oid).One(&u)
	if err!=nil{
		return u,err
	}
	return u,nil
}

func AlreadyLoggedIn(r *http.Request) (bool,error){
	c,err:=r.Cookie("session")
	if err!=nil{
		return false,errors.New("404. Cookie Not Found.")
	}
	s:=models.Session{}
	err=config.GetMongoDbSession().DB("demo").C("sessions").FindId(bson.ObjectIdHex(c.Value)).One(&s)
	if err!=nil{
		return false,errors.New("404. Session Not Found.")
	}
	u:=models.User{}
	err=config.GetMongoDbSession().DB("demo").C("users").FindId(s.UserId).One(&u)
	if err!=nil{
		return false,errors.New("404. User Not Found.")
	}
	return true,nil
}

