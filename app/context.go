package app

import "encoding/json"
import "net/http"

type Context struct {
	W    http.ResponseWriter
	R    *http.Request
	DB   *Database
	segs []string
	data data
}

type data map[string]interface{}
type Middleware func(*Context) bool
type Handler func(*Context)

func (c *Context) Set(k string, v interface{}) {
	c.data[k] = v
}

func (c *Context) Get(k string) (interface{}, bool) {
	v, ok := c.data[k]

	return v, ok
}

func (c *Context) Segs() []string {
	return c.segs
}

func (c *Context) Seg(n int) string {
	if len(c.segs) >= n {
		return c.segs[n-1]
	}

	return ""
}

func (c *Context) JSON(v interface{}, s ...int) {
	var code int

	if len(s) > 0 {
		code = s[0]
	} else {
		code = 200
	}

	b, _ := json.Marshal(v)
	c.W.Header().Set("Content-Type", "application/json")
	c.W.WriteHeader(code)
	c.W.Write(b)
}

func (c *Context) Error(s int) {
	c.JSON(Error{codes[s], s}, s)
}
