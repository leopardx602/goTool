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
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

// jwt secret key
var jwtSecret = []byte("secret")

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		// validate request body
		var body struct {
			Account  string
			Password string
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// check account and password is correct
		if body.Account == "username" && body.Password == "password" {
			now := time.Now()
			jwtId := body.Account + strconv.FormatInt(now.Unix(), 10)
			role := "Member"

			// set claims and sign
			claims := Claims{
				Account: body.Account,
				Role:    role,
				StandardClaims: jwt.StandardClaims{
					Audience:  body.Account,
					ExpiresAt: now.Add(20 * time.Second).Unix(),
					Id:        jwtId,
					IssuedAt:  now.Unix(),
					Issuer:    "ginJWT",
					NotBefore: now.Add(10 * time.Second).Unix(),
					Subject:   body.Account,
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

			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}

		// incorrect account or password
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	})

	// protected member router
	authorized := router.Group("/")
	authorized.Use(AuthRequired)
	{
		authorized.GET("/member/profile", func(c *gin.Context) {
			if c.MustGet("account") == "username" && c.MustGet("role") == "Member" {
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
	auth := c.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]

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
		fmt.Println("account:", claims.Account)
		fmt.Println("role:", claims.Role)
		c.Set("account", claims.Account)
		c.Set("role", claims.Role)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
