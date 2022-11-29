package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leopardx602/golang/google-oauth/model"
	"github.com/pkg/errors"
)

const (
	GoogleSecretKey = ""
	GoogleClientID  = ""
	RedirectURL     = "http://localhost:5000/api/oauth/google/login"
	Scope           = "https://www.googleapis.com/auth/userinfo.profile"
)

var (
	Users = map[string]User{}
)

type User struct {
	GoogleToken *model.GoogleToken
	UserInfo    *model.GoogleUserInfo
}

func main() {
	router := gin.Default()

	router.GET("/success", SuccessHandler)
	router.GET("/logout", func(c *gin.Context) { c.String(http.StatusOK, "logout") })
	router.GET("/api/oauth/google/url", GoogleAccess)
	router.GET("/api/oauth/google/login", GoogleLogin)
	router.Run(":5000")
}

func SuccessHandler(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	fmt.Println("success token:", token)
	c.String(http.StatusOK, "ok")
}

func GoogleAccess(c *gin.Context) {
	c.String(http.StatusOK, oauthURL())
}

func oauthURL() string {
	u := "https://accounts.google.com/o/oauth2/auth?client_id=%s&response_type=code&scope=%s&redirect_uri=%s"
	return fmt.Sprintf(u, GoogleClientID, Scope, RedirectURL)
}

func GoogleLogin(c *gin.Context) {
	code := c.Query("code")

	token, err := accessToken(code)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/logout")
		return
	}

	userInfo, err := getGoogleUserInfo(token.Token)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/logout")
		return
	}

	user := User{GoogleToken: token, UserInfo: userInfo}
	Users[userInfo.ID] = user
	fmt.Printf("id: %v, name: %v/n", userInfo.ID, userInfo.Name)

	c.SetCookie("token", token.Token, 3600, "/", c.Request.URL.Hostname(), false, true)
	c.Redirect(http.StatusFound, "/success")
}

func accessToken(code string) (token *model.GoogleToken, err error) {
	uri := "https://www.googleapis.com/oauth2/v4/token"
	data := url.Values{"code": {code}, "client_id": {GoogleClientID}, "client_secret": {GoogleSecretKey}, "grant_type": {"authorization_code"}, "redirect_uri": {RedirectURL}}
	body := strings.NewReader(data.Encode())

	resp, err := http.Post(uri, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("status code:%v, body:%v", resp.StatusCode, b)
	}

	token = &model.GoogleToken{}
	if err := json.Unmarshal(b, &token); err != nil {
		return token, err
	}
	return token, nil
}

func getGoogleUserInfo(token string) (userInfo *model.GoogleUserInfo, err error) {
	u := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("status code:%v, body:%v", resp.StatusCode, b)
	}

	userInfo = &model.GoogleUserInfo{}
	if err := json.Unmarshal(b, &userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
