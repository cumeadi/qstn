package models

import (
	"api/utils/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Model interface {
	Unique() bson.M
	Collection() string
}

type Hook interface {
	Model
	BeforeSave()
	AfterSave()
	BeforeCreate()
	AfterCreate()
}

func Cursor(m Model) *mgo.Collection {
	return db.DB.C(m.Collection())
}

func Find(m Model, q interface{}) *mgo.Query {
	return Cursor(m).Find(q)
}

func Update(m Hook) {
	m.BeforeSave()
	Cursor(m).Update(m.Unique(), m)
}

func Insert(m Hook) error {
	m.BeforeCreate()
	m.BeforeSave()
	if err := Cursor(m).Insert(m); err != nil {
		return err
	}
	m.AfterCreate()
	m.AfterSave()
	return nil
}
