package routes

import (
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/user/signup", controllers.userSignup).Methods("POST")
	r.HandleFunc("/user/login", controllers.userLogin).Methods("POST")
	r.HandleFunc("/user/getallusers", controllers.getAllusers).Methods("POST")

	r.HandleFunc("/owner/signup", controllers.ownerSignup).Methods("POST")
	r.HandleFunc("/owner/login", controllers.ownerLogin).Methods("POST")
	r.HandleFunc("/owner/getallusers", controllers.getAllowners).Methods("POST")

	r.HandleFunc("/admin/signup", controllers.adminSignup).Methods("POST")
	r.HandleFunc("/admin/login", controllers.adminLogin).Methods("POST")
	r.HandleFunc("/admin/getallusers", controllers.getAlladmins).Methods("POST")

}
