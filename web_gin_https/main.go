// $ go run /usr/local/go/src/crypto/tls/generate_cert.go --host="localhost"
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	router.Use(TlsHandler(":8000"))
	go router.RunTLS(":8000", "cert.pem", "key.pem")

	router.Run(":5000")

}

func TlsHandler(port string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     port,
		})
		if err := secureMiddleware.Process(c.Writer, c.Request); err != nil {
			fmt.Println(err)
			return
		}
		c.Next()
	}
}
