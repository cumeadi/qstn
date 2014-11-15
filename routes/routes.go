package routes

import (
	"github.com/bmizerany/pat"
	"github.com/codegangsta/negroni"
	"net/http"
)

func Map(m *pat.PatternServeMux, n *negroni.Negroni) {
	m.Get("/posts", http.HandlerFunc(postsGet))
	m.Post("/posts", http.HandlerFunc(postsCreate))
	m.Get("/posts/:id", http.HandlerFunc(postsShow))

	m.Get("/session", http.HandlerFunc(sessionGet))
}
