package main

import "github.com/daryl/qstn/app"
import "github.com/daryl/qstn/api"
import "github.com/daryl/qstn/models"
import "github.com/gorilla/websocket"
import "gopkg.in/mgo.v2/bson"
import "net/http"
import "time"
import "os"

const public = "public"

var hubs = map[string]*Hub{}

var grader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func main() {
	mux := app.New()

	mux.Use(static)
	mux.Use(buster)

	mux.Add("/", catch, filter)
	mux.Add("/q/", entries, filter)
	mux.Add("/s/", socket)

	mux.Listen(":7777")
}

func catch(c *app.Context) {
	c.Error(404)
}

func entries(c *app.Context) {
	segs := c.Segs()
	size := len(segs)

	switch {
	case size > 2:
		c.Error(404)
		return
	case size == 2:
		entriesID(c)
		return
	}

	switch c.R.Method {
	case "POST":
		entriesNew(c)
		return
	}

	c.Error(404)
}

func entriesID(c *app.Context) {
	if c.Seg(2) == "random" {
		entriesRandom(c)
		return
	}

	switch c.R.Method {
	case "GET":
		entriesGet(c)
		return
	}

	c.Error(404)
}

func entriesGet(c *app.Context) {
	id := c.Seg(2)

	s, entry := api.EntryGet(c, id)

	if s != 200 {
		c.Error(s)
		return
	}

	c.JSON(entry)
}

func entriesNew(c *app.Context) {
	s, entry := api.EntryPost(c)

	if s != 201 {
		c.Error(s)
		return
	}

	c.JSON(entry)
}

func entriesRandom(c *app.Context) {
	s, entry := api.EntryRand(c)

	if s != 200 {
		c.Error(s)
		return
	}

	c.JSON(entry)
}

func socket(c *app.Context) {
	ws, err := grader.Upgrade(c.W, c.R, nil)
	defer ws.Close()

	if err != nil {
		return
	}

	slug := c.Seg(2)
	coll := c.DB.C("entries")
	var entry models.Entry

	if _, ok := hubs[slug]; !ok {
		hubs[slug] = NewHub()
	}

	hub := hubs[slug]
	cli := hub.Add(ws)

	go cli.Ping(30 * time.Second)

	for {
		if err = ws.ReadJSON(&entry); err != nil {
			return
		}

		coll.Update(bson.M{
			"slug": slug,
		}, bson.M{
			"$set": entry,
		})

		for ws, _ := range hub.clients {
			ws.WriteJSON(entry)
		}
	}

	if len(hub.clients) == 0 {
		delete(hubs, slug)
	}

	hub.Remove(ws)
}

func filter(c *app.Context) bool {
	q := c.R.Header.Get("X-QSTN")

	if q == "" {
		http.ServeFile(c.W, c.R, "index.html")
		return false
	}

	return true
}

func buster(c *app.Context) bool {
	c.W.Header().Set("Cache-Control", "no-store")
	return true
}

func static(c *app.Context) bool {
	file := public + c.R.URL.Path

	s, err := os.Stat(file)

	if err == nil && !s.IsDir() {
		http.ServeFile(c.W, c.R, file)
		return false
	}

	return true
}
