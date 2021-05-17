package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Color string        `json:"color" bson:"color"`
}
