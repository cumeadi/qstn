package routes

import (
	"fmt"
	"github.com/daryl/sketchy-api/models"
	"github.com/daryl/sketchy-api/utils"
	"github.com/daryl/sketchy-api/utils/aws"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

func postsGet(w http.ResponseWriter, r *http.Request) {
	var post *models.Post
	var posts []*models.Post

	models.Find(post, nil).All(&posts)

	utils.JSON(w, posts)
}

func postsShow(w http.ResponseWriter, r *http.Request) {
	var post *models.Post

	models.Find(post, bson.M{
		"slug": r.URL.Query().Get(":id"),
	}).One(&post)

	utils.JSON(w, post)
}

func postsCreate(w http.ResponseWriter, r *http.Request) {
	// Parse form (25MB-ish)
	r.ParseMultipartForm((1 << 20) * 24)
	// Parse form values
	ff := r.MultipartForm.File
	fv := r.MultipartForm.Value
	id := bson.NewObjectId()
	// New Post instance
	p := &models.Post{
		Id:      id,
		Title:   fv["title"][0],
		Desc:    fv["desc"][0],
		Images:  []string{},
		Private: false,
	}

	// Convert tags to array
	if _, ok := fv["tags"]; ok {
		tags := strings.Split(fv["tags"][0], ",")
		for idx, t := range tags {
			tags[idx] = strings.TrimSpace(t)
		}
		p.Tags = tags
	}

	// Is it private?
	if _, yes := fv["private"]; yes {
		p.Private = true
	}

	// Create thumbnail
	for _, header := range ff["thumb"] {
		p.Thumb = header.Filename
		aws.PutImage(id.Hex(), header)
	}

	// Create images
	for _, header := range ff["images"] {
		p.Images = append(p.Images, header.Filename)
		aws.PutImage(id.Hex(), header)
	}

	// Create file
	for _, header := range ff["file"] {
		p.File = header.Filename
		aws.PutFile(id.Hex(), header)
	}

	fmt.Println(p)

	models.Insert(p)

	utils.JSON(w, p)
}

// func postsUpdate(w http.ResponseWriter, h http.Request) {
// 	// Parse form (25MB-ish)
// 	r.ParseMultipartForm((1 << 20) * 24)
// 	// Parse form values
// 	ff := r.MultipartForm.File
// 	fv := r.MultipartForm.Value
//
// 	var post *models.Post
// 	models.Find(post, bson.M{
// 		"_id": fv["_id"],
// 	}).One(&post)
//
// 	// Is it private?
// 	if _, yes := fv["private"]; yes {
// 		post.Private = true
// 	}
//
// 	for _, header := range ff["thumb"] {
// 		// New image, delete old one
// 		aws.PutImage(id.Hex(), header)
// 	}
//
// 	for _, header := range ff["images"] {
// 		// New image, delete old one
// 		aws.PutImage(id.Hex(), header)
// 	}
// }
