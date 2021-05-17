package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	UserId bson.ObjectId `json:"userId" bson:"userId"`
}
