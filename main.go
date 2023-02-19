package main

import (
	"net/http"

	"github.com/John-Santa/go-rest/db"
	"github.com/John-Santa/go-rest/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
