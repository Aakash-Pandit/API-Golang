package users

import (
	"encoding/json"
	"myapp/core"
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
	db_error := db.Find(&user, "id = ?", params["id"])
	if db_error.Error != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(map[string]string{"detail": "User Not Found"})
		return
	}

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

	validation_error := user.Validate()
	if validation_error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(validation_error)
		return
	}

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
	db_error := db.Find(&user, "id = ?", params["id"])
	if db_error.Error != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(map[string]string{"detail": "User Not Found"})
		return
	}
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
	db_error := db.Find(&user, "id = ?", params["id"])
	if db_error.Error != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(map[string]string{"detail": "User Not Found"})
		return
	}

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	validation_error := user.Validate()
	if validation_error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(validation_error)
		return
	}

	user.Modified = time.Now()
	db.Save(&user)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&user)
}

func SignIn(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var authentication UserAuthentication
	err := json.NewDecoder(request.Body).Decode(&authentication)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	validation_error := authentication.Validate()
	if validation_error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(validation_error)
		return
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var user User
	db_error := db.Find(&user, "email = ?", authentication.Email)
	if db_error.Error != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(map[string]string{"detail": "Email id Not Found"})
		return
	}

	if user.Password != authentication.Password {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]string{"detail": "Incorrect Password"})
		return
	}

	token, token_error := core.CreateToken(user.ID)
	if token_error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]string{"detail": token_error.Error()})
		return
	}

	token_obj := UserToken{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}

	json.NewEncoder(response).Encode(token_obj)
}
