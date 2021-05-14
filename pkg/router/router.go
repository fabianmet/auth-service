package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer(m *mux.Router) {
	log.Printf("Starting the server on %d\n", 8080)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", 8080), handlers.LoggingHandler(os.Stdout, m))
	if err != nil {
		log.Println(err)
	}
}

func NewRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	addCommonMiddleware(r)
	addCommonRoutes(r)
	printRoutes(r)

	return r, nil
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
func addCommonRoutes(m *mux.Router) {
	m.HandleFunc("/ping", pingHandler).Methods("GET")
	m.HandleFunc("/whoami", whoAmIHandler).Methods("GET")
	m.HandleFunc("/iam", iAmHandler).Methods("GET")
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
