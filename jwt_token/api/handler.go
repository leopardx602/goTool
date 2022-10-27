package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var param LoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else if _, ok := users[param.Username]; !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid username"))
		return
	} else if users[param.Username] != param.Password {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	token, err := Login(&param)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
