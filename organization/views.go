package organization

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllOrganizations(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var organizations []Organization

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Find(&organizations)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(organizations)
}

func GetOrganization(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var organization Organization
	db.Find(&organization, "id = ?", params["id"])

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(organization)
}

func CreateOrganization(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var organization Organization

	err := json.NewDecoder(request.Body).Decode(&organization)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	organization.Created = time.Now()
	organization.Modified = time.Now()

	db.Create(&organization)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(organization)
}

func DeleteOrganization(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var organization Organization
	db.Find(&organization, "id = ?", params["id"])
	db.Delete(&organization)

	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(map[string]string{})
}

func UpdateOrganization(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var organization Organization
	db.Find(&organization, "id = ?", params["id"])

	err = json.NewDecoder(request.Body).Decode(&organization)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}
	organization.Modified = time.Now()
	db.Save(&organization)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&organization)
}
