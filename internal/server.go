package internal

import (
	"context"
	"github.com/bwjson/api/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, handler *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:        cfg.Address,
			IdleTimeout: cfg.IdleTimeout,
			ReadTimeout: cfg.Timeout,
			Handler:     handler,
		},
	}
}

func (s *Server) Run() {
	s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
