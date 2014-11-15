package json

import (
	"net/http"
)

type JSON struct{}

func New() *JSON {
	return &JSON{}
}

func (c *JSON) ServeHTTP(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	n(w, r)
}
