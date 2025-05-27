package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*mux.Router
	shutdown chan os.Signal
	mw       []Middleware
}

func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleWare(mw, handler)
	handler = wrapMiddleWare(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		v := Values{
			TraceID: uuid.NewString(),
			Now:     time.Now().UTC(),
		}

		ctx := context.WithValue(r.Context(), key, &v)

		if err := handler(ctx, w, r); err != nil {
			fmt.Println(err)
			return
		}
	}

	a.HandleFunc(path, h).Methods(method)
}

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		Router:   mux.NewRouter(),
		shutdown: shutdown,
		mw:       mw,
	}
}
