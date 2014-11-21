package main

import (
	"github.com/codegangsta/negroni"
	"github.com/daryl/skatchy/middle/cors"
	"github.com/daryl/skatchy/middle/json"
	"github.com/daryl/skatchy/routes"
	"github.com/daryl/zeus"
	"os"
)

func main() {
	m := zeus.New()
	n := negroni.New()
	p := os.Getenv("PORT")
	// Middleware
	n.Use(cors.New())
	n.Use(json.New())
	// Make routes
	routes.Map(m, n)
	// Run server
	n.UseHandler(m)
	n.Run(":" + p)
}
