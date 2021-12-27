package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// custom claims
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// jwt secret key
var jwtSecret = []byte("secret")

var users = map[string]string{
	"username": "password",
	"user0001": "12345678",
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(c *gin.Context) {
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost:8080", false, true)
		c.String(200, "OK")
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})

	router.POST("/login", func(c *gin.Context) {
		// validate request body
		var body struct {
			Username string
			Password string
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Check the username.
		if _, ok := users[body.Username]; !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user",
			})
			return
		}

		// Check the password.
		if users[body.Username] != body.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// Generate the token
		now := time.Now()
		jwtId := body.Username + strconv.FormatInt(now.Unix(), 10)
		role := "Member"

		// set claims and sign
		claims := Claims{
			Username: body.Username,
			Role:     role,
			StandardClaims: jwt.StandardClaims{
				Audience:  body.Username,
				ExpiresAt: now.Add(3600 * time.Second).Unix(),
				Id:        jwtId,
				IssuedAt:  now.Unix(),
				Issuer:    "ginJWT",
				NotBefore: now.Add(5 * time.Second).Unix(),
				Subject:   body.Username,
			},
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		//c.SetCookie("token_123", token, 3600, "/", "localhost", false, true)
		fmt.Println(token)
		c.SetCookie("token", token, 3600, "/", "localhost:8080", false, true)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		//c.Redirect(http.StatusOK, "/member/profile")
	})

	// protected member router
	authorized := router.Group("/")
	authorized.Use(AuthRequired)
	{
		authorized.GET("/member/profile", func(c *gin.Context) {
			if c.MustGet("username") == "username" && c.MustGet("role") == "Member" {
				c.JSON(http.StatusOK, gin.H{
					"name":  "username",
					"age":   23,
					"hobby": "music",
				})
				return
			}

			c.JSON(http.StatusNotFound, gin.H{
				"error": "can not find the record",
			})
		})
	}
	router.Run()
}

// validate JWT
func AuthRequired(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Fail to get token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	fmt.Println(token)

	// auth := c.GetHeader("Authorization")
	// data := strings.Split(auth, "Bearer ")
	// if len(data) != 2 {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Unauthorized",
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// token := strings.Split(auth, "Bearer ")[1]

	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		fmt.Println("username:", claims.Username)
		fmt.Println("role:", claims.Role)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
