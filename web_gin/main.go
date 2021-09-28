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

	router.GET("/name", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"defaultName": "Chen"})
	})

	router.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"key1": "value1",
			"key2": "value2",
		})
	})

	router.POST("/service", func(ctx *gin.Context) {
		data := ctx.PostForm("userName")
		fmt.Println(data)
		ctx.JSON(200, gin.H{"status": "OK"})
	})

	router.GET("/img/:filename", func(ctx *gin.Context) {
		fileName := ctx.Param("filename")
		fmt.Println(fileName)
		ctx.File("./static/" + fileName)
	})

	router.Run(":5000")
}
