package aws

import (
	"bytes"
	"github.com/joho/godotenv"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

var con *s3.S3
var buk *s3.Bucket
var sizes = [4]uint{
	400,
	300,
	200,
	100,
}

func init() {
	godotenv.Load()
	a := os.Getenv("AWSKEY")
	s := os.Getenv("AWSSEC")
	k := os.Getenv("AWSBUK")
	con = s3.New(aws.Auth{a, s}, aws.EUWest)
	buk = con.Bucket(k)
}

func PutImage(id string, file *multipart.FileHeader) {
	open, _ := file.Open()
	bytz, _ := ioutil.ReadAll(open)

	path := filePath(id, file.Filename)
	kind := file.Header["Content-Type"][0]

	for _, size := range sizes {
		s := appendToFile(path, size)
		// Scale each size for retina
		for _, dpr := range [3]uint{3, 2, 1} {
			p := appendToFile(s, "@"+strconv.Itoa(int(dpr))+"x")
			b := resizeImage(bytz, size*dpr, kind)
			PutRaw(p, b, kind)
		}
	}
}

func PutFile(id string, file *multipart.FileHeader) {
	open, _ := file.Open()
	bytz, _ := ioutil.ReadAll(open)

	path := filePath(id, file.Filename)
	kind := file.Header["Content-Type"][0]

	PutRaw(path, bytz, kind)
}

func PutRaw(path string, bytz []byte, kind string) {
	err := buk.Put(path, bytz, kind, s3.BucketOwnerFull)
	if err != nil {
		panic(err)
	}
}

func resizeImage(bytz []byte, width uint, kind string) []byte {
	buf := new(bytes.Buffer)
	dec, _, _ := image.Decode(bytes.NewReader(bytz))
	rsz := resize.Resize(width, 0, dec, resize.Lanczos3)
	switch kind {
	case "image/jpeg":
		jpeg.Encode(buf, rsz, &jpeg.Options{100})
	case "image/png":
		png.Encode(buf, rsz)
	}
	return buf.Bytes()
}

func appendToFile(file string, val interface{}) string {
	parts := strings.Split(file, ".")
	switch val.(type) {
	case uint:
		parts[0] += strconv.Itoa(int(val.(uint)))
	case string:
		parts[0] += val.(string)
	}
	return strings.Join(parts, ".")
}

func filePath(id, file string) string {
	return id + "/" + file
}
