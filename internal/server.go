package internal

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, handler *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:        cfg.HttpSrv.Address,
			IdleTimeout: cfg.HttpSrv.IdleTimeout,
			ReadTimeout: cfg.HttpSrv.Timeout,
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
