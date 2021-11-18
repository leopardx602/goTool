package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

//var store = sessions.NewFilesystemStore("./", securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

var store *sessions.CookieStore

func init() {
	//store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
	store = sessions.NewCookieStore([]byte("key1234"))
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	fmt.Println("session server run on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["auth"] = true
	session.Values["name"] = "chen"
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "logged in")
}

func home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if auth, ok := session.Values["auth"].(bool); !ok || !auth {
		http.Error(w, "unauthorizeed", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, session.Values["name"])
}
