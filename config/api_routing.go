package config

import (
	"myapp/patients"
	"myapp/userapp"

	"github.com/gorilla/mux"
)

func UserApiRouting(router *mux.Router) {
	userapp.CreateDummyUsers()
	patients.CreateDummyPatients()

	router.HandleFunc("/api/v1/users", userapp.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", userapp.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/", userapp.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", userapp.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/v1/users/{id}", userapp.DeleteUser).Methods("DELETE")

	router.HandleFunc("/api/v1/patients", patients.GetAllPatients).Methods("GET")
	router.HandleFunc("/api/v1/patients/{id}", patients.GetPatient).Methods("GET")
	router.HandleFunc("/api/v1/patients/", patients.CreatePatient).Methods("POST")
	router.HandleFunc("/api/v1/patients/{id}", patients.UpdatePatient).Methods("PATCH")
	router.HandleFunc("/api/v1/patients/{id}", patients.DeletePatient).Methods("DELETE")
}
