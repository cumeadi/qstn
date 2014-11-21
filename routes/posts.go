package routes

import (
	"fmt"
	m "github.com/daryl/skatchy/models"
	"github.com/daryl/skatchy/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func postsGet(w http.ResponseWriter, r *http.Request) {
	var post *m.Post
	var posts []*m.Post

	m.Find(post, nil).All(&posts)

	for i, _ := range posts {
		posts[i].ToJSON()
	}

	utils.JSON(w, posts)
}

func postsShow(w http.ResponseWriter, r *http.Request) {
	var post *m.Post

	m.Find(post, bson.M{
		"slug": r.URL.Query().Get("id"),
	}).One(&post)

	if post != nil {
		post.ToJSON()
	}

	fmt.Println(post)

	utils.JSON(w, post)
}

func postsCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm((1 << 20) * 24)

	//ff := r.MultipartForm.File
	fv := r.MultipartForm.Value
	id := bson.NewObjectId()

	//hex := id.Hex()

	p := &m.Post{
		Id:      id,
		Title:   fv["title"][0],
		Desc:    fv["desc"][0],
		Tags:    fv["tags"],
		Private: false,
	}

	//	files := []*multipart.FileHeader{}
	//
	//	for _, header := range ff["thumb"] {
	//		p.Thumb.File = header.Filename
	//		p.Thumb.Pref = hex
	//		// Upload to S3.
	//		p.Thumb.Put(header)
	//	}

	w.WriteHeader(201)
	utils.JSON(w, p)
}
