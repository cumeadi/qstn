package assets

import (
	"mime/multipart"
)

type Asset interface {
	Put(f *multipart.FileHeader) error
	Get() string
	Del() error
}
