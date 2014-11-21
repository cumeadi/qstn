package aws

import (
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"os"
)

var (
	bucket *s3.Bucket
)

func init() {
	bucket = s3.New(aws.Auth{
		os.Getenv("AWS_ACCESS"),
		os.Getenv("AWS_SECRET"),
	}, aws.EUWest).Bucket(
		os.Getenv("AWS_BUCKET"),
	)
}

func Get(path string) string {
	return bucket.URL(path)
}

func Put(path string, bytz []byte, kind string) error {
	return bucket.Put(path, bytz, kind, s3.BucketOwnerFull)
}
