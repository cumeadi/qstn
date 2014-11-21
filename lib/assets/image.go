package assets

import (
	"bytes"
	"fmt"
	"github.com/daryl/skatchy/lib/aws"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"
)

var (
	scales = [3]int{
		3,
		2,
		1,
	}
	sizes = [2]int{
		400,
		200,
	}
)

type Image struct {
	File string          `json:"file"`
	Size map[string]size `json:"size" bson:"-"`
	Pref string          `json:"-"`
	err  chan error      `json:"-"`
}

type size struct {
	Scale map[string]scale `json:"scale"`
}

type scale struct {
	Path string `json:"path"`
}

func (i Image) Put(file *multipart.FileHeader) error {
	open, _ := file.Open()
	fata, _ := ioutil.ReadAll(open)
	kind := file.Header["Content-Type"][0]
	path := i.Pref + "/" + i.File

	// Total amount of images
	total := len(sizes) * len(scales)
	// Error channel
	i.err = make(chan error, total)

	// Loop through sizes
	for _, size := range sizes {
		go func(size int) {
			_path := appendToFile(path, size)
			for _, scale := range scales {
				go func(scale int) {
					n := strconv.Itoa(scale)
					p := appendToFile(_path, "@"+n+"x")
					b := resizeImage(fata, size*scale, kind)
					i.err <- aws.Put(p, b, kind)
				}(scale)
			}
		}(size)
	}

	// Block / Error handler
	for y := 0; y < total; y++ {
		if err := <-i.err; err != nil {
			return err
		}
	}

	return nil
}

func (i Image) Get() {
	fmt.Println(i.err)

	path := i.Pref + "/" + i.File

	i.Size = make(map[string]size)

	for _, size := range sizes {
		fmt.Println(size)
	}

	url := aws.Get(path)

	fmt.Println(url)
}

// Resize an image to a certain width (keeping ratio).
func resizeImage(bytz []byte, width int, kind string) []byte {
	buf := new(bytes.Buffer)
	raw, _, _ := image.Decode(bytes.NewReader(bytz))
	rsz := resize.Resize(uint(width), 0, raw, resize.Lanczos3)
	switch kind {
	case "image/jpeg":
		jpeg.Encode(buf, rsz, &jpeg.Options{100})
	case "image/png":
		png.Encode(buf, rsz)
	}
	return buf.Bytes()
}

// Append value to filename.
func appendToFile(file string, val interface{}) string {
	parts := strings.Split(file, ".")
	switch val.(type) {
	case int:
		parts[0] += strconv.Itoa(val.(int))
	case string:
		parts[0] += val.(string)
	}
	return strings.Join(parts, ".")
}
