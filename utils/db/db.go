package db

import (
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"os"
)

var DB *mgo.Database

func init() {
	godotenv.Load()
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	DB = s.DB(os.Getenv("DB"))
}
