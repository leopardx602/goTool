package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Validate JWT
func AuthCookie(c *gin.Context) {
	// Get the token from the cookie or the authorization
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	claims, err := authenticate(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("unexpected claim error"))
		return
	}
	c.Set("username", claims.Username)
	c.Set("role", claims.Role)
	c.Next()
}

func AuthBearerToken(c *gin.Context) {
	// Get the token from the authorization of header
	auth := c.GetHeader("Authorization")
	data := strings.Split(auth, "Bearer ")
	if len(data) != 2 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("error in header Authorization"))
		return
	}
	token := strings.Split(auth, "Bearer ")[1]

	claims, err := authenticate(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("unexpected claim error"))
		return
	}
	c.Set("username", claims.Username)
	c.Set("role", claims.Role)
	c.Next()
}

// Parse and validate token
func authenticate(token string) (claims *Claims, err error) {
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
		return nil, errors.New(message)
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, errors.New("unexpected claim error")
}
