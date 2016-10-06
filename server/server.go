package server

import (
	"trom/auth"
	"trom/gateway"
)

type Server struct {
}

func New(g *gateway.Gateway, auth *auth.Auth) *Server {

	return new(Server)
}

func (s *Server) Start() {

}
