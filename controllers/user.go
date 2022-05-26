package controllers

import (
	"fmt"
	"net/http"
)

func userSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}

func getAllusers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getallusers")
}
