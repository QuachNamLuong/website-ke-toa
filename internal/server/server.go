package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/quachnamluong/website-ke-toa/internal/app"
	"github.com/quachnamluong/website-ke-toa/internal/router"
)

type Server struct {
	httpServer *http.Server
	app *app.App
}

func New(a *app.App, addr string) *Server {
	
	s := &Server{app: a}

	s.httpServer = &http.Server{
		Addr: addr,
		Handler: router.New(a),
	}

	return s
}

func (s *Server) Start() error {
	s.app.Logger.Info("starting server", "addr", s.httpServer.Addr)

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}

// Shutdown gracefully stops the server, waiting up to the given timeout
// for in-flight requests to finish.
func (s *Server) Shutdown(ctx context.Context) error {
	s.app.Logger.Info("shutting down server")
	return s.httpServer.Shutdown(ctx)
}