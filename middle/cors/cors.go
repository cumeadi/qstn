package cors

import (
	"net/http"
)

type CORS struct{}

func New() *CORS {
	return &CORS{}
}

func (c *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5001")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Allow")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	n(w, r)
}
