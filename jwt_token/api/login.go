package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var (
	jwtSecret = []byte("secret")
	users     = map[string]string{
		"username": "password",
		"user0001": "12345678",
	}
)

func Login(param *LoginParam) (token string, err error) {
	// Check the username and password.
	if _, ok := users[param.Username]; !ok {
		return "", errors.New("Invalid username")
	} else if users[param.Username] != param.Password {
		return "", errors.New("Invalid password")
	}

	// Generate the token
	now := time.Now()
	jwtId := param.Username + strconv.FormatInt(now.Unix(), 10)
	role := "Member"

	// Set claims and sign
	claims := Claims{
		Username: param.Username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			Audience:  param.Username,
			ExpiresAt: now.Add(3600 * time.Second).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(5 * time.Second).Unix(),
			Subject:   param.Username,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}
