package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/imayrus/e-commerce/database"
	"github.com/imayrus/e-commerce/models"
)

func ownerSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var owner models.Owner
	// Unmarshal
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		log.Fatal(err, "unable to decode")
		return
	}

	// Validate the struct
	validate := validator.New()
	validationErr := validate.Struct(owner)
	if validationErr != nil {
		log.Fatal(err, "Validation err")
		return
	}

	temp := &models.Owner{}
	// Check for errors and duplicate usernames
	database.GetDB().Where("username = ?", owner.Username).First(temp)
	if temp.Username != "" {
		log.Fatal(err, "Username already exist")
	}

	database.GetDB().Where("email = ?", owner.Email).First(temp)
	if temp.Email != "" {
		log.Fatal(err, "email already exists")
		json.NewEncoder(w).Encode(err)
		return
	}

	// Create database and Encode
	database.GetDB().Create(owner)
	log.Println("Successfully signed up ")
	json.NewEncoder(w).Encode(owner)
}

func ownrLogin(w http.ResponseWriter, r *http.Request) {
	var owner models.Owner
	var foundowner models.Owner
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		log.Fatal(err, "unable to decode")
		return
	}

	database.GetDB().Where("email = ?", owner.Email).First(&foundowner)
	if foundowner.Email != "" {
		log.Fatal(err, "Email already exists")
		json.NewEncoder(w).Encode(err)
	}

	database.GetDB().Where("password = ?", owner.Password).First(&foundowner)
	if foundowner.Email != "" {
		log.Fatal(err, "Password already exists")
		json.NewEncoder(w).Encode(err)
	}
}

func getAllowners(w http.ResponseWriter, r *http.Request) {
	var owners models.Owner
	database.GetDB().Find(owners)
}
