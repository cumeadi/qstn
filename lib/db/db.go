package db

import (
	"gopkg.in/mgo.v2"
	"os"
)

var (
	DB *mgo.Database
)

func init() {
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	DB = s.DB(os.Getenv("DB"))
}
