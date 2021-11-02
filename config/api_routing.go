package config

import (
	"myapp/core"
	"myapp/organization"
	"myapp/patients"
	"myapp/users"

	"github.com/gorilla/mux"
)

func UserApiRouting(router *mux.Router) {

	router.HandleFunc("/api/v1/users", core.IsAuthorized(users.GetAllUsers)).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", core.IsAuthorized(users.GetUser)).Methods("GET")
	router.HandleFunc("/api/v1/users/", users.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", core.IsAuthorized(users.UpdateUser)).Methods("PATCH")
	router.HandleFunc("/api/v1/users/{id}", core.IsAuthorized(users.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/api/v1/users/signup", users.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/signin", users.SignIn).Methods("POST")

	router.HandleFunc("/api/v1/patients", core.IsAuthorized(patients.GetAllPatients)).Methods("GET")
	router.HandleFunc("/api/v1/patients/{id}", core.IsAuthorized(patients.GetPatient)).Methods("GET")
	router.HandleFunc("/api/v1/patients/", core.IsAuthorized(patients.CreatePatient)).Methods("POST")
	router.HandleFunc("/api/v1/patients/{id}", core.IsAuthorized(patients.UpdatePatient)).Methods("PATCH")
	router.HandleFunc("/api/v1/patients/{id}", core.IsAuthorized(patients.DeletePatient)).Methods("DELETE")

	router.HandleFunc("/api/v1/medicines", core.IsAuthorized(patients.GetAllMedicines)).Methods("GET")
	router.HandleFunc("/api/v1/medicines/", core.IsAuthorized(patients.CreateMedicine)).Methods("POST")
	router.HandleFunc("/api/v1/medicines/{id}", core.IsAuthorized(patients.GetMedicine)).Methods("GET")
	router.HandleFunc("/api/v1/medicines/{id}", core.IsAuthorized(patients.DeleteMedicine)).Methods("DELETE")
	router.HandleFunc("/api/v1/medicines/{id}", core.IsAuthorized(patients.UpdateMedicine)).Methods("PATCH")

	router.HandleFunc("/api/v1/organizations", core.IsAuthorized(organization.GetAllOrganizations)).Methods("GET")
	router.HandleFunc("/api/v1/organizations/", core.IsAuthorized(organization.CreateOrganization)).Methods("POST")
	router.HandleFunc("/api/v1/organizations/{id}", core.IsAuthorized(organization.GetOrganization)).Methods("GET")
	router.HandleFunc("/api/v1/organizations/{id}", core.IsAuthorized(organization.DeleteOrganization)).Methods("DELETE")
	router.HandleFunc("/api/v1/organizations/{id}", core.IsAuthorized(organization.UpdateOrganization)).Methods("PATCH")
}
