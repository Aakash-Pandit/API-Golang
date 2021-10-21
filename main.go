package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var portNumber int = 8080

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Contact   string    `json:"contact"`
}

var users []User

func Home(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"detail": "This is API Home Page of version V1",
	}
	response.WriteHeader((200))
	json.NewEncoder(response).Encode(data)
}

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader((200))
	json.NewEncoder(response).Encode(users)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)
	for _, item := range users {
		if item.ID.String() == params["id"] {
			response.WriteHeader((200))
			json.NewEncoder(response).Encode(item)
			return
		}
	}
	response.WriteHeader((200))
	json.NewEncoder(response).Encode(map[string]string{})
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}
	user.BeforeCreate()
	users = append(users, user)
	response.WriteHeader((201))
	json.NewEncoder(response).Encode(user)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)

	for index, item := range users {
		if item.ID.String() == params["id"] {
			users = append(users[:index], users[index+1:]...)
			response.WriteHeader((204))
			json.NewEncoder(response).Encode(map[string]string{})
			return
		}
	}
	response.WriteHeader((400))
	json.NewEncoder(response).Encode(map[string]string{"details": "Bad Request"})
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range users {

		var user User
		if item.ID.String() == params["id"] {
			users = append(users[:index], users[index+1:]...)
			err := json.NewDecoder(request.Body).Decode(&user)
			if err != nil {
				response.WriteHeader((400))
				json.NewEncoder(response).Encode(err)
				return
			}
			response.WriteHeader((200))
			id, _ := uuid.Parse(params["id"])
			user.ID = id
			users = append(users, user)
			json.NewEncoder(response).Encode(user)
			return
		}
	}

	response.WriteHeader((400))
	json.NewEncoder(response).Encode(map[string]string{})
}

func (user *User) BeforeCreate() User {
	(*user).ID = uuid.New()
	return *user
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	users = append(users, User{ID: uuid.New(), FirstName: "Aakash", LastName: "Pandit", Email: "aakashpandit366@gmail.com", Contact: "8698410175"})
	users = append(users, User{ID: uuid.New(), FirstName: "Siddhesh", LastName: "Pandit", Email: "Sid@gmail.com", Contact: "1234543210"})
	users = append(users, User{ID: uuid.New(), FirstName: "Rasika", LastName: "Pandit", Email: "ras@gmail.com", Contact: "1234543210"})

	router.HandleFunc("/api/v1", Home).Methods("GET")
	router.HandleFunc("/api/v1/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/", CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/v1/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe((":" + strconv.Itoa(portNumber)), router))
}
