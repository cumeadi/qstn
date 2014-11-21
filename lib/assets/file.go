package assets

import (
	"github.com/daryl/skatchy/lib/aws"
	"io/ioutil"
	"mime/multipart"
)

type File struct {
	File string     `json:"file"`
	Pref string     `json:"pref"`
	Path string     `json:"path" bson:"-"`
	err  chan error `json:"-"`
}

func (f File) Put(file *multipart.FileHeader) error {
	open, _ := file.Open()
	fata, _ := ioutil.ReadAll(open)
	kind := file.Header["Content-Type"][0]
	path := f.Pref + "/" + f.File

	return aws.Put(path, fata, kind)
}

func (f File) Get() string {
	path := f.Pref + "/" + f.File
	return aws.Get(path)
}
