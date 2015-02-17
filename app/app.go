package app

import "github.com/daryl/qstn/utils/str"
import "net/http"

type App struct {
	mux *http.ServeMux
	mw  []Middleware
}

func New() *App {
	return &App{
		http.NewServeMux(),
		[]Middleware{},
	}
}

func (a *App) Listen(port string) {
	http.ListenAndServe(port, a.mux)
}

func (a *App) Use(mw ...Middleware) {
	a.mw = append(a.mw, mw...)
}

func (a *App) Add(p string, h Handler, mw ...Middleware) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		db := copyDB()
		defer db.Close()

		ctx := &Context{
			w, r, db,
			str.Split(str.Trim(r.URL.Path, "/"), "/"),
			data{},
		}

		for _, mw := range a.mw {
			if ok := mw(ctx); !ok {
				return
			}
		}

		for _, mw := range mw {
			if ok := mw(ctx); !ok {
				return
			}
		}

		h(ctx)
	}

	a.mux.HandleFunc(p, fn)
}
