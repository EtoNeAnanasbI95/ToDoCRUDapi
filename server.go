package ToDoCRUD

import (
	"context"
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler, cfg *config.Config) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%v", cfg.Api.Port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.Api.Timeout,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
