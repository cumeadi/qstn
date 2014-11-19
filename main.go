package main

import (
	"github.com/codegangsta/negroni"
	"github.com/daryl/sketchy-api/middle/cors"
	"github.com/daryl/sketchy-api/middle/json"
	"github.com/daryl/sketchy-api/routes"
	"github.com/daryl/zeus"
)

func main() {
	m := zeus.New()
	n := negroni.New()
	// Middleware
	n.Use(cors.New())
	n.Use(json.New())
	// Make routes
	routes.Map(m, n)
	// Run server
	n.UseHandler(m)
	n.Run(":5000")
}
