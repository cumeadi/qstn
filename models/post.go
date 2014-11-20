package models

import (
	"github.com/daryl/sketchy-api/utils/aws"
	"gopkg.in/mgo.v2/bson"
	"regexp"
	"strings"
	"time"
)

type Image struct {
	File string `json:"file"`
	Size map[string]struct {
		Dpr map[string]string `json:"dpr"`
	} `json:"size"`
}

type Post struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Slug      string        `json:"slug"`
	Title     string        `json:"title"`
	Desc      string        `json:"desc"`
	Thumb     string        `json:"thumb"`
	ThumbURL  aws.Image     `json:"thumbURL" bson:"-"`
	File      string        `json:"file"`
	FileURL   string        `json:"fileURL" bson:"-"`
	Images    []string      `json:"images"`
	ImagesURL aws.Images    `json:"imagesURL" bson"-"`
	Tags      []string      `json:"tags"`
	Views     int           `json:"views"`
	Downloads int           `json:"downloads"`
	Private   bool          `json:"private"`
	Updated   time.Time     `json:"updated"`
	Made      time.Time     `json:"made"`
}

func (p *Post) AsJSON() {
	id := p.Id.Hex()

	p.ImagesURL = make(aws.Images)

	for _, image := range p.Images {
		urls := aws.GetImage(id, image)
		p.ImagesURL[image] = urls
	}

	if p.Thumb != "" {
		p.ThumbURL = aws.GetImage(id, p.Thumb)
	}

	if p.File != "" {
		p.FileURL = aws.GetFile(id, p.File)
	}
}

func (p *Post) Unique() bson.M {
	return bson.M{"_id": p.Id}
}

func (_ *Post) Collection() string {
	return "posts"
}

func (p *Post) BeforeCreate() {
	p.Views = 0
	p.Downloads = 0
	// Created time
	p.Made = time.Now()
	// If the slug isn't
	// already set, create
	// one using the title.
	if p.Slug != "" {
		p.setSlug()
	}
}

func (p *Post) BeforeSave() {
	// Updated time
	p.Updated = time.Now()
}

func (_ *Post) AfterCreate() {}
func (_ *Post) AfterSave()   {}

func (p *Post) setSlug() {
	r := regexp.MustCompile("[^\\w-]{1,}")
	s := r.ReplaceAllString(p.Title, "-")
	p.Slug = strings.ToLower(string(s))
}
