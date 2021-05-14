package main

import (
	"github.com/fabianmet/auth-service/pkg/router"
)

func main() {
	r, err := router.NewRouter()
	if err != nil {
		panic(err)
	}

	router.StartServer(r)
}
