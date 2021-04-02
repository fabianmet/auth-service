package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	r := gin.Default()

	err := getRoutes(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func getRoutes(r *gin.Engine) error {
	auth := r.Group("/auth")
	addCommonRoutes(auth)
	addGoogleRoutes(auth)

	return nil
}

// addCommonRoutes creates the common routes. Exposing public keys and
func addCommonRoutes(rg *gin.RouterGroup) {
	rg.GET("/ping", pingHandler)
	rg.GET("/pubkey", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
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

//pingHandler because im playing around
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
