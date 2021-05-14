package main

import (
	"github.com/fabianmet/auth-service/pkg/router"
	"github.com/fabianmet/auth-service/pkg/server"
)

func main() {

	s := server.NewServer()

	r := router.NewRouter(s)

	r.StartServer()
}
