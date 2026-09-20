package handler

import (
	"net/http"

	"github.com/quachnamluong/website-ke-toa/internal/app"
	viewApp "github.com/quachnamluong/website-ke-toa/view/app"
)

type AppHandler struct {
	app *app.App
}

func NewAppHandler(a *app.App) *HomeHandler {
	return &HomeHandler{app: a}
}

func (h *AppHandler) Show(w http.ResponseWriter, r *http.Request) {
	if err := viewApp.Index().Render(r.Context(), w); err != nil {
		h.app.Logger.Error("render home failed", "err", err) // <- panics if h.app or h.app.Logger is nil
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}