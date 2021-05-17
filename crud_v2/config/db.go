package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

var DB *sql.DB

func GetMongoDbSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}

func init()  {
	var err error
	DB,err=sql.Open("postgres", "postgres://hhz:hhz123@localhost/demo?sslmode=disable")
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
	if err=DB.Ping();err!=nil{
		panic(err)
	}
	//fmt.Println("You Connected to Postgres...")
}


