package api

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler, port string) *Server {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(5) * time.Second,
	}
	return &Server{server: server}
}

func (h *Server) Run() error {
	return h.server.ListenAndServe()
}

func (h *Server) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
