package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var portNumber int = 8080

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Contact   string `json:"contact"`
}

var users []User

func Home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"detail": "This is API Home Page of version V1",
	}
	writer.WriteHeader((200))
	json.NewEncoder(writer).Encode(data)
}

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader((200))
	json.NewEncoder(writer).Encode(users)
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)
	for _, item := range users {
		if item.ID == params["id"] {
			writer.WriteHeader((200))
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	writer.WriteHeader((200))
	json.NewEncoder(writer).Encode(map[string]string{})
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "appication/json")
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	users = append(users, user)
	writer.WriteHeader((201))
	json.NewEncoder(writer).Encode(user)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)

	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			writer.WriteHeader((204))
			json.NewEncoder(writer).Encode(map[string]string{})
			return
		}
	}
	writer.WriteHeader((400))
	json.NewEncoder(writer).Encode(map[string]string{"details": "Bad Request"})
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range users {

		var user User
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			err := json.NewDecoder(request.Body).Decode(&user)
			if err != nil {
				writer.WriteHeader((400))
				json.NewEncoder(writer).Encode(err)
				return
			}
			writer.WriteHeader((200))
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(writer).Encode(user)
			return
		}
	}

	writer.WriteHeader((400))
	json.NewEncoder(writer).Encode(map[string]string{})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	users = append(users, User{ID: "1", FirstName: "Aakash", LastName: "Pandit", Email: "aakashpandit366@gmail.com", Contact: "8698410175"})
	users = append(users, User{ID: "2", FirstName: "Siddhesh", LastName: "Pandit", Email: "Sid@gmail.com", Contact: "1234543210"})
	users = append(users, User{ID: "3", FirstName: "Rasika", LastName: "Pandit", Email: "ras@gmail.com", Contact: "1234543210"})

	router.HandleFunc("/api/v1", Home).Methods("GET")
	router.HandleFunc("/api/v1/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/", CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/v1/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe((":" + strconv.Itoa(portNumber)), router))
}
