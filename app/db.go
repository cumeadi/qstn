package app

import "gopkg.in/mgo.v2"
import "os"

type Database struct {
	*mgo.Database
}

var ds *mgo.Session
var db string

func init() {
	var err error

	ds, err = mgo.Dial(os.Getenv("MONGOURI"))

	if err != nil {
		panic(err)
	}

	db = os.Getenv("MONGODB")

	ds.DB(db).C("entries").EnsureIndex(mgo.Index{
		Unique: true,
		Key: []string{
			"slug",
		},
	})
}

func copyDB() *Database {
	return &Database{
		ds.Copy().DB(db),
	}
}

func (d *Database) Close() {
	d.Session.Close()
}
