package controllers

import (
	"fmt"
	"net/http"
)

func adminSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}

func getAlladmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getallusers")
}
