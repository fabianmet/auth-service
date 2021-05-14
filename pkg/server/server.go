package server

import (
	"github.com/fabianmet/auth-service/pkg/config"
	"github.com/fabianmet/auth-service/pkg/jwt"
)

type Server struct {
	Key    *jwt.RsaKey
	Config *config.Config
}

func NewServer() *Server {
	config := config.NewConfig()
	key := jwt.NewRsaKey()

	return &Server{
		Key:    key,
		Config: config,
	}
}
