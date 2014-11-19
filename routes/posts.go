package routes

import (
	"github.com/daryl/sketchy-api/models"
	"github.com/daryl/sketchy-api/utils"
	"github.com/daryl/sketchy-api/utils/aws"
	"gopkg.in/mgo.v2/bson"
	"mime/multipart"
	"net/http"
)

type assets map[string]*multipart.FileHeader

func postsGet(w http.ResponseWriter, r *http.Request) {
	var post *models.Post
	var posts []*models.Post

	models.Find(post, nil).All(&posts)

	utils.JSON(w, posts)
}

func postsShow(w http.ResponseWriter, r *http.Request) {
	var post *models.Post

	models.Find(post, bson.M{
		"slug": r.URL.Query().Get("id"),
	}).One(&post)

	// TODO: Change
	if inc := r.URL.Query().Get("inc"); inc != "" {
		switch inc {
		case "downloads":
			post.Downloads++
		case "views":
			post.Views++
		}
		models.Update(post)
	}

	utils.JSON(w, post)
}

func postsCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm((1 << 20) * 24)

	ff := r.MultipartForm.File
	fv := r.MultipartForm.Value
	id := bson.NewObjectId()

	p := &models.Post{
		Id:      id,
		Title:   fv["title"][0],
		Desc:    fv["desc"][0],
		Tags:    fv["tags"],
		Images:  []string{},
		Private: false,
	}

	if val, ok := ff["thumb"]; ok {
		p.Thumb = val[0].Filename
	}

	if val, ok := ff["file"]; ok {
		p.File = val[0].Filename
	}

	if _, ok := fv["private"]; ok {
		p.Private = true
	}

	for _, val := range ff["images"] {
		p.Images = append(p.Images, val.Filename)
	}

	files := []*multipart.FileHeader{}

	for _, items := range ff {
		files = append(files, items...)
	}

	aws.PutFiles(id.Hex(), files)

	models.Insert(p)

	utils.JSON(w, p)
}
