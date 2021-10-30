package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var users []User

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Find(&users)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var user User
	db.Find(&user, "id = ?", params["id"])

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	var user User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	user.Created = time.Now()
	user.Modified = time.Now()

	db.Create(&user)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(user)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var user User
	db.Find(&user, "id = ?", params["id"])
	db.Delete(&user)

	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(map[string]string{})
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var user User
	db.Find(&user, "id = ?", params["id"])

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	user.Modified = time.Now()
	db.Save(&user)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&user)
}
