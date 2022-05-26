package controllers

import (
	"fmt"
	"net/http"
)

func ownerSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
}

func ownrLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}

func getAllowners(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getallusers")
}
