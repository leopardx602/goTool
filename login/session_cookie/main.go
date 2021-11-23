package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store *sessions.CookieStore
var userData = map[string]string{}

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
	userData["admin"], _ = HashAndSalt("admin")
	userData["chen"], _ = HashAndSalt("1234")
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

// Check login status
func isLogin(w http.ResponseWriter, request *http.Request) bool {
	session, err := store.Get(request, "session-name")
	if err != nil {
		log.Println(err)
		return false
	}

	if auth, ok := session.Values["auth"].(bool); !ok || !auth {
		return false
	}
	return true
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["auth"] = nil
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "logged out.")
}

func login(response http.ResponseWriter, request *http.Request) {
	// get
	if request.Method == http.MethodGet {
		if ok := isLogin(response, request); ok {
			http.Redirect(response, request, "/", http.StatusFound)
			return
		}
		fmt.Fprintln(response, `
			<h1>Login</h1>
			<form method="post" action="/login">
				<label for="name">User name</label>
				<input type="text" id="name" name="name">
				<label for="password">Password</label>
				<input type="password" id="password" name="password">
				<button type="submit">Login</button>
			</form>
			`)
		return
	}

	// post
	name := request.FormValue("name")
	pass := request.FormValue("password")
	if name == "" || pass == "" || !ComparePasswords(userData[name], pass) {
		http.Redirect(response, request, "/", http.StatusFound)
	}

	// Set session informations
	session, err := store.Get(request, "session-name")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["auth"] = true
	session.Values["name"] = "chen"
	if err := session.Save(request, response); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(response, request, "/", http.StatusFound)
}

func home(w http.ResponseWriter, r *http.Request) {
	if ok := isLogin(w, r); !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Hello %v", session.Values["name"])
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
