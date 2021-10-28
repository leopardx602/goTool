package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Discount bool   `json:"discount"`
}

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
		// m := map[string]string{"status": "ok"}
		// j, _ := json.Marshal(m)
		// c.Data(http.StatusOK, "application/json", j)
	})

	router.GET("/struct", func(ctx *gin.Context) {
		var product []Product
		product = append(product, Product{"apple", 1000, true})
		product = append(product, Product{"orange", 2000, false})
		data, err := json.Marshal(product)
		if err != nil {
			fmt.Println(err)
		}
		ctx.Data(200, "application/json", data)

	})

	// post form
	router.POST("/service", func(ctx *gin.Context) {
		data := ctx.PostForm("userName")
		fmt.Println(data)
		ctx.JSON(200, gin.H{"status": "OK"})
	})

	// post struct
	router.POST("/product", func(ctx *gin.Context) {
		var product Product
		ctx.BindJSON(&product)
		fmt.Println(product)
		ctx.JSON(200, gin.H{"status": "OK"})
	})

	// post map
	router.POST("/product2", func(ctx *gin.Context) {
		product := make(map[string]interface{})
		ctx.BindJSON(&product)
		fmt.Println(product)
		ctx.JSON(200, gin.H{"status": "OK"})
	})

	router.GET("/img/:filename", func(ctx *gin.Context) {
		fileName := ctx.Param("filename")
		fmt.Println(fileName)
		ctx.File("./static/" + fileName)
	})

	router.DELETE("product", func(ctx *gin.Context) {
		product := make(map[string]interface{})
		ctx.BindJSON(&product)
		fmt.Println(product)
		ctx.JSON(200, gin.H{"status": "OK"})
	})

	router.Run(":5000")
}
