package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mobile/config"
	"net/http"
	"time"
)

const (
	_shutdownTimout = 3 * time.Second
)

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

func NewHttpServer(gin *gin.Engine, cfg *config.HttpService) *Server {
	httpServer := http.Server{
		Addr:         cfg.Address,
		Handler:      gin,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	server := &Server{
		server:          &httpServer,
		shutdownTimeout: _shutdownTimout,
	}

	return server
}

func (s *Server) Start() error {
	const fn = "httpServer.Start"

	err := s.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("%v: %s", fn, err)
	}
	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
