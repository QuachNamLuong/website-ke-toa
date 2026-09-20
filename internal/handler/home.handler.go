package handler

import (
	"net/http"

	"github.com/quachnamluong/website-ke-toa/internal/app"
	"github.com/quachnamluong/website-ke-toa/view"
)

type HomeHandler struct {
	app *app.App
}

func NewHomeHandler(a *app.App) *HomeHandler {
	return &HomeHandler{app: a}
}

func (h *HomeHandler) Show(w http.ResponseWriter, r *http.Request) {
	if err := view.Home().Render(r.Context(), w); err != nil {
		h.app.Logger.Error("render home failed", "err", err) // <- panics if h.app or h.app.Logger is nil
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}