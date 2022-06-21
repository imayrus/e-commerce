package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imayrus/e-commerce/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}
