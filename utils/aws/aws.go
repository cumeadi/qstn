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

type Image map[string]map[string]string

type Images map[string]Image

var bucket *s3.Bucket

var sizes = [4]int{
	400,
	300,
	200,
	100,
}

var dpr = [3]int{
	3,
	2,
	1,
}

func init() {
	godotenv.Load()
	a := os.Getenv("AWSKEY")
	s := os.Getenv("AWSSEC")
	k := os.Getenv("AWSBUK")
	bucket = s3.New(aws.Auth{
		a, s,
	}, aws.EUWest).Bucket(k)
}

func GetImage(id string, filename string) Image {
	out := make(Image)

	path := filePath(id, filename)

	for _, size := range sizes {
		s := strconv.Itoa(size)
		// Create size
		out[s] = map[string]string{}
		// Append size to path.
		_path := appendToFile(path, size)
		// Loop through @3x, @2x, @1x.
		for _, dpr := range dpr {
			n := strconv.Itoa(dpr)
			x := "@" + n + "x"
			p := appendToFile(_path, x)
			out[s][x] = p
		}
	}

	return out
}

func GetFile(id string, filename string) string {
	return filePath(id, filename)
}

// Put multiple files to S3 at once.
func PutFiles(id string, files []*multipart.FileHeader) error {
	nf := len(files)
	// Success channel.
	sc := make(chan bool, nf)
	// Error channel.
	ec := make(chan error, nf)

	// Look through files.
	for _, file := range files {
		// Send out a separate goroutine
		// for each file. Much fast. Wow.
		go func(file *multipart.FileHeader) {
			// Namespace path with ObjectID.
			path := filePath(id, file.Filename)
			// Images require resizing.
			if isImage(file) {
				putImage(path, file, sc, ec)
				return
			}
			// Regular file.
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

// Resize images then PUT each to S3.
func putImage(path string, file *multipart.FileHeader, sc chan bool, ec chan error) {
	open, _ := file.Open()
	bytz, _ := ioutil.ReadAll(open)

	kind := file.Header["Content-Type"][0]

	var wg sync.WaitGroup
	// For all the sizes
	// multiply by 3 to
	// support <= @3x.
	wg.Add(len(sizes) * 3)

	// Loop through sizes.
	for _, size := range sizes {
		// New goroutine.
		go func(size int) {
			// Append size to path.
			_path := appendToFile(path, size)
			// Loop through @3x, @2x, @1x.
			for _, dpr := range dpr {
				// New goroutine.
				go func(dpr int) {
					n := strconv.Itoa(dpr)
					p := appendToFile(_path, "@"+n+"x")
					// Resize image. Like a boss.
					b := resizeImage(bytz, size*dpr, kind)
					// Send to S3.
					putRaw(p, b, kind, ec)
					wg.Done()
				}(dpr)
			}
			wg.Done()
		}(size)
	}

	wg.Wait()

	sc <- true
}

// PUT a file to S3.
func putFile(path string, file *multipart.FileHeader, sc chan bool, ec chan error) {
	open, _ := file.Open()
	bytz, _ := ioutil.ReadAll(open)

	kind := file.Header["Content-Type"][0]

	// Send to S3.
	putRaw(path, bytz, kind, ec)

	sc <- true
}

// PUT bytes to S3.
func putRaw(path string, bytz []byte, kind string, ec chan error) {
	err := bucket.Put(path, bytz, kind, s3.BucketOwnerFull)
	// Error?
	if err != nil {
		ec <- err
	}
	// Success
	fmt.Println(path)
}

// Resize an image to a certain width (keeping ratio).
func resizeImage(bytz []byte, width int, kind string) []byte {
	buf := new(bytes.Buffer)
	dec, _, _ := image.Decode(bytes.NewReader(bytz))
	rsz := resize.Resize(uint(width), 0, dec, resize.Lanczos3)
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

// Create the filepath from ObjectID.
func filePath(id, file string) string {
	return id + "/" + file
}

// Test whether a file is an image.
func isImage(file *multipart.FileHeader) bool {
	switch file.Header["Content-Type"][0] {
	case "image/jpeg", "image/png":
		return true
	default:
		return false
	}
}
