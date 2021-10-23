package config

import (
	"myapp/userapp"

	"github.com/gorilla/mux"
)

func UserApiRouting(router *mux.Router) {
	userapp.CreateDummyUsers()

	router.HandleFunc("/api/v1/users", userapp.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", userapp.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/", userapp.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", userapp.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/v1/users/{id}", userapp.DeleteUser).Methods("DELETE")
}
