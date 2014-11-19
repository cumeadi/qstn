package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/daryl/zeus"
)

func Map(m *zeus.Mux, n *negroni.Negroni) {
	m.GET("/posts", postsGet)
	m.POST("/posts", postsCreate)
	m.GET("/posts/:id", postsShow)

	m.GET("/users", usersGet)
	m.POST("/users", usersCreate)
	m.GET("/users/:id", usersShow)

	m.GET("/session", sessionGet)
}
