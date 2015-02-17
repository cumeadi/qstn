package models

import "gopkg.in/mgo.v2/bson"

type Entry struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Slug     string        `bson:",omitempty" json:"slug"`
	Question string        `bson:",omitempty" json:"question"`
	Options  []option      `bson:",omitempty" json:"options"`
}

type option struct {
	Option string `bson:"option" json:"option"`
	Votes  int    `bson:"votes" json:"votes"`
}
