package api

import (
	"encoding/json"
	"github.com/daryl/qstn/app"
	"github.com/daryl/qstn/models"
	"github.com/daryl/qstn/utils/num"
	"github.com/daryl/qstn/utils/str"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func EntryGet(c *app.Context, slug string) (int, models.Entry) {
	var entry models.Entry

	err := c.DB.C("entries").Find(bson.M{
		"slug": slug,
	}).One(&entry)

	if err == mgo.ErrNotFound {
		return 404, entry
	}

	return 200, entry
}

func EntryRand(c *app.Context) (int, models.Entry) {
	var entry models.Entry

	coll := c.DB.C("entries")
	size, err := coll.Count()

	if err != nil {
		return 500, entry
	}

	err = coll.Find(nil).Skip(
		num.RandBetween(0, size),
	).One(&entry)

	if err == mgo.ErrNotFound {
		return 404, entry
	}

	return 200, entry
}

func EntryPost(c *app.Context) (int, models.Entry) {
	var entry models.Entry
	json.NewDecoder(c.R.Body).Decode(&entry)
	entry.ID = bson.NewObjectId()

	if entry.Question == "" || len(entry.Options) < 2 {
		return 400, entry
	}

	coll := c.DB.C("entries")

	for {
		entry.Slug = str.Rand(8)

		has, _ := coll.Find(bson.M{
			"slug": entry.Slug,
		}).Count()

		if has < 1 {
			break
		}
	}

	if err := coll.Insert(entry); err != nil {
		return 500, entry
	}

	return 201, entry
}
