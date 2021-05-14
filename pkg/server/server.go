package server

import (
	"github.com/fabianmet/auth-service/pkg/config"
	"github.com/fabianmet/auth-service/pkg/jwt"
	"github.com/fabianmet/auth-service/pkg/router"
	"github.com/gorilla/mux"
)

type Server struct {
	Key    *jwt.RsaKey
	Router *mux.Router
	Config *config.Config
}

func NewServer() *Server {
	config := config.NewConfig()
	key := jwt.NewRsaKey()
	router, err := router.NewRouter()
	if err != nil {
		panic(err)
	}

	return &Server{
		Key:    key,
		Config: config,
		Router: router,
	}

}
