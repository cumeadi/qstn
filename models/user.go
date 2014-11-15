package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password []byte        `json:"-"`
	Token    string        `json:"token"`
}

func (u *User) Unique() bson.M {
	return bson.M{"_id": u.Id}
}

func (u *User) Collection() string {
	return "users"
}
