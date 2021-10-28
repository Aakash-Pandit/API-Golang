package config

import (
	"myapp/patients"
	"myapp/users"

	"github.com/gorilla/mux"
)

func UserApiRouting(router *mux.Router) {
	users.CreateDummyUsers()
	patients.CreateDummyPatients()

	router.HandleFunc("/api/v1/users", users.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", users.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/", users.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", users.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/v1/users/{id}", users.DeleteUser).Methods("DELETE")

	router.HandleFunc("/api/v1/patients", patients.GetAllPatients).Methods("GET")
	router.HandleFunc("/api/v1/patients/{id}", patients.GetPatient).Methods("GET")
	router.HandleFunc("/api/v1/patients/", patients.CreatePatient).Methods("POST")
	router.HandleFunc("/api/v1/patients/{id}", patients.UpdatePatient).Methods("PATCH")
	router.HandleFunc("/api/v1/patients/{id}", patients.DeletePatient).Methods("DELETE")

	router.HandleFunc("/api/v1/medicines", patients.GetAllMedicines).Methods("GET")
	router.HandleFunc("/api/v1/medicines/", patients.CreateMedicine).Methods("POST")
	router.HandleFunc("/api/v1/medicines/{id}", patients.GetMedicine).Methods("GET")
	router.HandleFunc("/api/v1/medicines/{id}", patients.DeleteMedicine).Methods("DELETE")
	router.HandleFunc("/api/v1/medicines/{id}", patients.UpdateMedicine).Methods("PATCH")
}
