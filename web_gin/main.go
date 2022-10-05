package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Discount bool   `json:"discount"`
	Soldout  *bool  `json:"soldout"`
	Email    string `json:"email" binding:"required" validate:"email"`
}

var (
	validate = validator.New()
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/name", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"defaultName": "Chen"})
	})

	router.GET("/string", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world!")
	})

	router.GET("/slice", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, []string{"a", "b", "c"})
	})

	router.GET("/map", func(ctx *gin.Context) {
		data := map[string]string{"key1": "value1", "key2": "value2"}
		ctx.JSON(http.StatusOK, data)
		//ctx.JSON(http.StatusOK, gin.H{"key1": "value1", "key2": "value2"})
	})

	router.GET("/struct", func(ctx *gin.Context) {
		var product struct {
			Name  string
			Price int
		}
		product.Name = "iphone"
		product.Price = 1000
		ctx.JSON(http.StatusOK, product)
	})

	// post form
	router.POST("/service", func(ctx *gin.Context) {
		data := ctx.PostForm("userName")
		fmt.Println(data)
		ctx.String(http.StatusOK, "ok")
	})

	// post struct
	router.POST("/product", func(ctx *gin.Context) {
		var product Product
		if err := ctx.BindJSON(&product); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		} else if err := validate.Struct(product); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		fmt.Printf("%+v\n", product)

		if product.Soldout == nil {
		} else {
			fmt.Println(*product.Soldout)
		}
		ctx.String(http.StatusOK, "ok")
	})

	// post map
	router.POST("/product2", func(ctx *gin.Context) {
		product := make(map[string]interface{})
		ctx.BindJSON(&product)
		fmt.Println(product)
		ctx.String(http.StatusOK, "ok")
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
		ctx.String(http.StatusOK, "ok")
	})

	// css javascript
	router.Static("/static", "./static")

	router.Run(":5000")
}
