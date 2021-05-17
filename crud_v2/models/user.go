package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	Password  string        `json:"password" bson:"password"`
}
