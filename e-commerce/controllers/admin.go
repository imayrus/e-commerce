package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/imayrus/e-commerce/database"
	"github.com/imayrus/e-commerce/models"
)

func adminSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var admin models.Admin
	// Unmarshal
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		log.Fatal(err, "unable to decode")
		return
	}

	// Validate the struct
	validate := validator.New()
	validationErr := validate.Struct(admin)
	if validationErr != nil {
		log.Fatal(err, "Validation err")
		return
	}

	temp := &models.Admin{}
	// Check for errors and duplicate adminnames
	database.GetDB().Where("adminname = ?", admin.Username).First(temp)
	if temp.Username != "" {
		log.Fatal(err, "Adminname already exist")
	}

	database.GetDB().Where("email = ?", admin.Email).First(temp)
	if temp.Email != "" {
		log.Fatal(err, "email already exists")
		json.NewEncoder(w).Encode(err)
		return
	}

	// Create database and Encode
	database.GetDB().Create(admin)
	log.Println("Successfully signed up ")
	json.NewEncoder(w).Encode(admin)
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	var foundadmin models.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		log.Fatal(err, "unable to decode")
		return
	}

	database.GetDB().Where("email = ?", admin.Email).First(&foundadmin)
	if foundadmin.Email != "" {
		log.Fatal(err, "Email already exists")
		json.NewEncoder(w).Encode(err)
	}

	database.GetDB().Where("password = ?", admin.Password).First(&foundadmin)
	if foundadmin.Email != "" {
		log.Fatal(err, "Password already exists")
		json.NewEncoder(w).Encode(err)
	}
}

func getAlladmin(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	database.GetDB().Find(admin)
}
