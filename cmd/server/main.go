package main

import (
	"github.com/fabianmet/auth-service/pkg/jwt"
	"github.com/fabianmet/auth-service/pkg/router"
)

func main() {

	key := jwt.NewRsaKey()

	r, err := router.NewRouter()
	if err != nil {
		panic(err)
	}

	router.StartServer(r)
}
