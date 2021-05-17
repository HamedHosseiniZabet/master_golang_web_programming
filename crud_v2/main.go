package main

import (
	"github.com/master_go_programming/crud_v2/config"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/master_go_programming/crud_v2/controllers"
)

func main() {
	r := httprouter.New()
	hc := controllers.NewHomeController(config.GetMongoDbSession())
	pc:=controllers.NewProductController(config.GetMongoDbSession())
	r.GET("/signup", hc.SignUp) // view must change
	r.POST("/signupProcess",hc.SignUpProcess)
	r.POST("/loginProcess",hc.LoginProcess)
	r.POST("/addProduct",pc.AddProduct)
	r.GET("/getProduct/:id",pc.GetProduct)
	r.PUT("/editProduct/:id",pc.EditProduct)
	r.GET("/getProducts",pc.GetProducts)
	r.DELETE("/deleteProduct/:id",pc.DeleteProduct)
	r.GET("/login", hc.Login)
	r.GET("/logout", hc.Logout)
	r.GET("/", hc.Home)
	r.ServeFiles("/assets/*filepath", http.Dir("./assets"))
	log.Fatal(http.ListenAndServe(":8080", r))
}

