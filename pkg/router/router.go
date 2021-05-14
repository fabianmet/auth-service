package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	r := gin.Default()

	err := getRoutes(r)
	if err != nil {
		return nil, err
	}

	store(r)

	return r, nil
}

func getRoutes(r *gin.Engine) error {
	auth := r.Group("/auth")
	addCommonRoutes(auth)
	addGoogleRoutes(auth)
	addLocalRoutes(auth)

	return nil
}

func store(r *gin.Engine) {
	store := cookie.NewStore([]byte("secretyolo"))
	r.Use(sessions.Sessions("mysession", store))
}

// addCommonRoutes creates the common routes. Exposing public keys and
func addCommonRoutes(rg *gin.RouterGroup) {
	rg.GET("/ping", pingHandler)

	// pubkey is the endpoint that will show the public key you can use to validate jwt tokens.
	rg.GET("/pubkey", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	rg.GET("/whoami", whoAmIHandler)
}

// addGoogleRoutes creates the google specific routes
func addGoogleRoutes(rg *gin.RouterGroup) {
	// TODO: no inline functions please, this is bad
	google := rg.Group("/google")

	google.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "google login")
	})

	google.GET("/callback", func(c *gin.Context) {
		c.JSON(http.StatusOK, "google callback")
	})

	google.GET("/refresh", func(c *gin.Context) {
		c.JSON(http.StatusOK, "google refresh")
	})
}

// addLocalRoutes creates the Local specific routes
func addLocalRoutes(rg *gin.RouterGroup) {
	local := rg.Group("/local")

	local.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "local login")
	})

	local.GET("/callback", func(c *gin.Context) {
		c.JSON(http.StatusOK, "local callback")
	})

	local.GET("/refresh", func(c *gin.Context) {
		c.JSON(http.StatusOK, "local refresh")
	})
}

//pingHandler because im playing around
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//whoAmIHandler returns information about who you are.
func whoAmIHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to implement",
	})
}

//iAmHandler returns information about who you are.
func iamHandler(c *gin.Context) {

}
