package session

import (
	"api/utils"
	"net/http"
	"time"
)

const (
	KEY = "go"
)

var store = map[string]*session{}

type session struct {
	id   string
	data map[string]interface{}
	time time.Time
}

func (s *session) Get(key string) interface{} {
	return s.data[key]
}

func (s *session) Set(key string, val interface{}) {
	s.data[key] = val
}

func Start(w http.ResponseWriter, r *http.Request) *session {
	cookie, err := r.Cookie(KEY)

	if err != nil {
		return new(w)
	}

	sesh, ok := store[cookie.Value]

	if !ok {
		return new(w)
	}

	if time.Now().After(sesh.time) {
		delete(store, sesh.id)
		return new(w)
	}

	return sesh
}

func new(w http.ResponseWriter) *session {
	id := utils.RandSeq(32)

	store[id] = &session{
		id,
		map[string]interface{}{},
		time.Now().AddDate(0, 1, 0),
	}

	http.SetCookie(w, &http.Cookie{
		Name:     KEY,
		Value:    id,
		HttpOnly: true,
	})

	return store[id]
}
