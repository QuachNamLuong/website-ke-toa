package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/quachnamluong/website-ke-toa/internal/app"
	"github.com/quachnamluong/website-ke-toa/internal/handler"
)

func New(a *app.App) http.Handler {
	r := chi.NewRouter()
	
	r.Get("/", handler.NewHomeHandler(a).Show)
	r.Get("/app", handler.NewAppHandler(a).Show)
	return r
}