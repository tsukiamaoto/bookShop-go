package server

import (
	"tsukiamaoto/bookShop-go/config"

	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(conf *config.Config, handler http.Handler) *Server {
	return &Server {
		httpServer: &http.Server{
			Addr: conf.ServerAddress,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}