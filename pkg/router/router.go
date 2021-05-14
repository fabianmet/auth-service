package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fabianmet/auth-service/pkg/server"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Router struct {
	Server *server.Server
	Muxer  *mux.Router
}

func (r *Router) StartServer() {
	log.Printf("Starting the server on %d\n", 8080)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", 8080), handlers.LoggingHandler(os.Stdout, r.Muxer))
	if err != nil {
		log.Println(err)
	}
}

func NewRouter(s *server.Server) *Router {
	r := mux.NewRouter()

	router := &Router{
		Server: s,
		Muxer:  r,
	}

	addCommonMiddleware(r)
	router.addCommonRoutes(r)
	printRoutes(r)

	return router
}

func addCommonMiddleware(m *mux.Router) {
	m.Use(mwAddContext)
}

func printRoutes(m *mux.Router) {
	m.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, _ := route.GetPathTemplate()
		log.Printf("Accepting on path %v\n", template)
		return nil
	})
}

// addCommonRoutes creates the common routes. Exposing public keys and
func (router *Router) addCommonRoutes(m *mux.Router) {
	m.HandleFunc("/ping", router.pingHandler).Methods("GET")
	m.HandleFunc("/whoami", router.whoAmIHandler).Methods("GET")
	m.HandleFunc("/iam", router.iAmHandler).Methods("GET")
}

// // addGoogleRoutes creates the google specific routes
// func addGoogleRoutes(rg *gin.RouterGroup) {
// 	// TODO: no inline functions please, this is bad
// 	google := rg.Group("/google")

// 	google.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "google login")
// 	})

// 	google.GET("/callback", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "google callback")
// 	})

// 	google.GET("/refresh", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "google refresh")
// 	})
// }

// // addLocalRoutes creates the Local specific routes
// func addLocalRoutes(rg *gin.RouterGroup) {
// 	local := rg.Group("/local")

// 	local.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "local login")
// 	})

// 	local.GET("/callback", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "local callback")
// 	})

// 	local.GET("/refresh", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "local refresh")
// 	})
// }
