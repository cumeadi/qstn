package aws

import (
	"bytes"
	"fmt"
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
	"sync"
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

func PutFiles(id string, files []*multipart.FileHeader) error {
	nf := len(files)
	// Success / Error chan
	sc := make(chan bool, nf)
	ec := make(chan error, nf)

	for _, file := range files {
		go func(file *multipart.FileHeader) {
			path := filePath(id, file.Filename)
			if isImage(file) {
				putImage(path, file, sc, ec)
				return
			}
			putFile(path, file, sc, ec)
		}(file)
	}

	for {
		select {
		case _ = <-sc:
			nf--
		case err := <-ec:
			return err
		}
		if nf == 0 {
			break
		}
	}

	return nil
}

func putImage(path string, file *multipart.FileHeader, sc chan bool, ec chan error) {
	open, _ := file.Open()
	bytz, _ := ioutil.ReadAll(open)

	kind := file.Header["Content-Type"][0]

	var wg sync.WaitGroup
	wg.Add(len(sizes))

	for _, size := range sizes {
		go func(size uint) {
			_path := appendToFile(path, size)
			for _, dpr := range [3]uint{3, 2, 1} {
				wg.Add(1)
				go func(dpr uint) {
					// Path to resized image and dpr (image200@2x)
					p := appendToFile(_path, "@"+strconv.Itoa(int(dpr))+"x")
					// Resize image
					b := resizeImage(bytz, size*dpr, kind)
					// Send to S3
					putRaw(p, b, kind, ec)
					// DONE
					wg.Done()
				}(dpr)
			}
			// DONE
			wg.Done()
		}(size)
	}

	wg.Wait()
	// Success
	sc <- true
}

func putFile(path string, file *multipart.FileHeader, sc chan bool, ec chan error) {
	open, _ := file.Open()
	bytz, _ := ioutil.ReadAll(open)

	kind := file.Header["Content-Type"][0]

	putRaw(path, bytz, kind, ec)

	// Success
	sc <- true
}

func putRaw(path string, bytz []byte, kind string, ec chan error) {
	err := buk.Put(path, bytz, kind, s3.BucketOwnerFull)
	// Error?
	if err != nil {
		ec <- err
	}
	// Success
	fmt.Println(path)
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

func isImage(file *multipart.FileHeader) bool {
	switch file.Header["Content-Type"][0] {
	case "image/jpeg", "image/png":
		return true
	default:
		return false
	}
}
