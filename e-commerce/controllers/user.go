package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/imayrus/e-commerce/database"
	"github.com/imayrus/e-commerce/models"
)

func userSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	// Unmarshal
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err, "unable to decode")
		return
	}

	// Validate the struct
	validate := validator.New()
	validationErr := validate.Struct(user)
	if validationErr != nil {
		log.Fatal(err, "Validation err")
		return
	}

	temp := &models.User{}
	// Check for errors and duplicate usernames
	database.GetDB().Where("username = ?", user.Username).First(temp)
	if temp.Username != "" {
		log.Fatal(err, "Username already exist")
	}

	database.GetDB().Where("email = ?", user.Email).First(temp)
	if temp.Email != "" {
		log.Fatal(err, "email already exists")
		json.NewEncoder(w).Encode(err)
		return
	}

	// Create database and Encode
	database.GetDB().Create(user)
	log.Println("Successfully signed up ")
	json.NewEncoder(w).Encode(user)

}

func userLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var founduser models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err, "unable to decode")
		return
	}

	database.GetDB().Where("email = ?", user.Email).First(&founduser)
	if founduser.Email != "" {
		log.Fatal(err, "Email already exists")
		json.NewEncoder(w).Encode(err)
	}

	database.GetDB().Where("password = ?", user.Password).First(&founduser)
	if founduser.Email != "" {
		log.Fatal(err, "Password already exists")
		json.NewEncoder(w).Encode(err)
	}

}

func getAllusers(w http.ResponseWriter, r *http.Request) {
	var users models.User
	database.GetDB().Find(users)
}
