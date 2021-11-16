package main

import (
	"fmt"

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

func main() {
	if pass := ComparePasswords(userData["chen"], "1234"); pass {
		fmt.Println("pass")
	} else {
		fmt.Println("deny")
	}

	if pass := ComparePasswords(userData["ting"], "0000"); pass {
		fmt.Println("pass")
	} else {
		fmt.Println("deny")
	}
}
