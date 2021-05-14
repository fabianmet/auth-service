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

//pingHandler because im playing around
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

//whoAmIHandler returns information about who you are.
func whoAmIHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("idtoken")

	// assume not logged in
	if err != nil {
		http.Redirect(w, r, "/iam", http.StatusFound)
		return
	}
	fmt.Println(cookie.Value)
}

//iAmHandler returns information about who you are.
func iAmHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "idtoken",
		Value:    "token goes here",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

// func store(r *gin.Engine) {
// 	store := cookie.NewStore([]byte("secretyolo"))
// 	r.Use(sessions.Sessions("mysession", store))
// }

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
