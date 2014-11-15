package models

import (
	"gopkg.in/mgo.v2/bson"
	"regexp"
	"strings"
	"time"
)

type Post struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Slug      string        `json:"slug"`
	Title     string        `json:"title"`
	Desc      string        `json:"desc"`
	Thumb     string        `json:"thumb"`
	File      string        `json:"file"`
	Images    []string      `json:"images"`
	Tags      []string      `json:"tags"`
	Views     int           `json:"views"`
	Downloads int           `json:"downloads"`
	Private   bool          `json:"private"`
	Made      time.Time     `json:"made"`
}

func (p *Post) Unique() bson.M {
	return bson.M{"_id": p.Id}
}

func (_ *Post) Collection() string {
	return "posts"
}

func (p *Post) BeforeCreate() {
	p.setSlug()
	p.Views = 0
	p.Made = time.Now()
	p.Downloads = 0
}

func (p *Post) setSlug() {
	r := regexp.MustCompile("[^\\w-]{1,}")
	s := r.ReplaceAllString(p.Title, "-")
	p.Slug = strings.ToLower(string(s))
}

func (_ *Post) AfterCreate() {}
func (_ *Post) BeforeSave()  {}
func (_ *Post) AfterSave()   {}
