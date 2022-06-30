package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "good")
}

func main() {
	router := gin.Default()
	router.GET("/healthcheck", HealthCheckHandler)

	router.Run()
}
