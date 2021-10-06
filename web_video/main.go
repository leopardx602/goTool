package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/static/:filename", func(ctx *gin.Context) {
		fileName := ctx.Param("filename")
		fmt.Println(fileName)
		ctx.File("./static/" + fileName)
	})
	router.Run(":5000")
}
