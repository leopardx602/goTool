package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

		// Set claims and sign
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

		fmt.Println(token)
		c.SetCookie("token", token, 3600, "/", "localhost:8080", false, true)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	// Authorized router
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

// Validate JWT
func AuthRequired(c *gin.Context) {
	// Get the token from the cookie or the authorization
	token, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Failed to get the token from the cookie")
		auth := c.GetHeader("Authorization")
		data := strings.Split(auth, "Bearer ")
		if len(data) != 2 {
			fmt.Println("Failed to get the token from the authorization")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		token = strings.Split(auth, "Bearer ")[1]
	}
	fmt.Println(token)

	// Parse and validate token
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
