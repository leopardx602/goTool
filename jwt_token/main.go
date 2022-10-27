package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leopardx602/golang/jwt_token/api"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.POST("/login", api.LoginHandler)

	// Authorized router
	authorized := router.Group("/content")
	authorized.Use(api.AuthCookie)
	authorized.GET("/hello", func(c *gin.Context) {
		username, ok := c.Get("username")
		if !ok {
			c.AbortWithError(http.StatusInternalServerError, errors.New("error in getting username"))
		}
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", username))
	})

	router.Run()
}
