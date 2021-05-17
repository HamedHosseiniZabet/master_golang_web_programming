package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/master_go_programming/crud_v2/config"
	"github.com/master_go_programming/crud_v2/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type ProductController struct {
	session *mgo.Session
}
func NewProductController(s *mgo.Session) *ProductController{
	return &ProductController{session: s}
}

func (pc ProductController) AddProduct(w http.ResponseWriter,r *http.Request,_ httprouter.Params)  {
	p:=models.Product{}
	p.Id=bson.NewObjectId()
	p.Name=r.FormValue("name")
	p.Color=r.FormValue("color")
	err:=pc.session.DB("demo").C("products").Insert(p)
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	if p.Name==""||p.Color==""{
		http.Error(w,http.StatusText(400),http.StatusBadRequest)
		return
	}
	// convert form values
	//f64, err := strconv.ParseFloat(p, 32)
	//if err != nil {
	//	http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
	//	return
	//}
	//bk.Price = float32(f64)
	_,err=config.DB.Exec("INSERT INTO products(name,color) VALUES($1,$2)",p.Name,p.Color)
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w,p)

}
func (pc ProductController) GetProducts(w http.ResponseWriter,r *http.Request, _ httprouter.Params)  {
	var p []models.Product
	err:=pc.session.DB("demo").C("products").Find(bson.M{}).All(&p)
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	pj,err:= json.Marshal(p)
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pj)
}
func (pc ProductController) DeleteProduct(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	p:=models.Product{}
	id:=ps.ByName("id")
	if !bson.IsObjectIdHex(id){
		http.Error(w,http.StatusText(404),http.StatusNotFound)
		return
	}
	err:=pc.session.DB("demo").C("products").FindId(bson.ObjectIdHex(id)).One(&p)
	if err!=nil{
		http.Error(w,http.StatusText(404),http.StatusNotFound)
		return
	}
	err=pc.session.DB("demo").C("products").RemoveId(bson.ObjectIdHex(id))
	if err!=nil{
		http.Error(w,http.StatusText(404),http.StatusNotFound)
		return
	}
	_,err=config.DB.Exec("DELETE FROM products WHERE name=$1",p.Name)
	if err!=nil{
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w,p)
}
func (pc ProductController) GetProduct(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	id := ps.ByName("id")

	if !bson.IsObjectIdHex(id) {
		http.Error(w,http.StatusText(404),http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	p := models.Product{}

	if err := pc.session.DB("demo").C("products").FindId(oid).One(&p); err != nil {
		http.Error(w,http.StatusText(404),http.StatusNotFound)
		return
	}

	pj, err := json.Marshal(p)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(pj)
}
func (pc ProductController) EditProduct(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	p:=models.Product{}
	ep:=models.Product{}
	id:=ps.ByName("id")
	n:=r.FormValue("name")
	c:=r.FormValue("color")
	if !bson.IsObjectIdHex(id){
		http.Error(w,http.StatusText(400),http.StatusBadRequest)
		return
	}
	if n==""||c==""{
		http.Error(w,http.StatusText(400),http.StatusBadRequest)
		return
	}
	ep.Id=bson.ObjectIdHex(id)
	ep.Name=n
	ep.Color=c
	err:=pc.session.DB("demo").C("products").FindId(bson.ObjectIdHex(id)).One(&p)
	if err!=nil{
		http.Error(w,http.StatusText(400),http.StatusBadRequest)
		return
	}
	err = pc.session.DB("demo").C("products").Update(bson.M{"_id":bson.ObjectIdHex(id)},&ep)
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	_,err=config.DB.Exec("UPDATE products SET name=$1, color=$2 WHERE name=$3",n,c,p.Name)
	if err!=nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
}

