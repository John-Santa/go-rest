package main

import (
	"net/http"

	"github.com/John-Santa/go-rest/db"
	"github.com/John-Santa/go-rest/models"
	"github.com/John-Santa/go-rest/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	router := mux.NewRouter()
	//Users routes
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")
	//Tasks routes
	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
