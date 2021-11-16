package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var userData = map[string]string{}

func init() {
	userData["admin"], _ = HashAndSalt("admin")
	userData["chen"], _ = HashAndSalt("1234")
	userData["ting"], _ = HashAndSalt("5678")
}

// Encrypt the password
func HashAndSalt(pwdStr string) (string, error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

// Authentication
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	if err := bcrypt.CompareHashAndPassword(byteHash, bytePwd); err != nil {
		return false
	}
	return true
}

// login handler
func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		if ComparePasswords(userData[name], pass) {
			redirectTarget = "/internal"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// index page
const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
}

// internal page
const internalPage = `<h1>Internal, %s</h1>`

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, internalPage, "hello world")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)
	router.HandleFunc("/login", loginHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
