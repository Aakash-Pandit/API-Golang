package main

import (
	"encoding/json"
	"log"
	"myapp/config"
	"myapp/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var portNumber int = 8080

func Home(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"detail": "This is API Home Page of version V1",
	}
	response.WriteHeader((200))
	json.NewEncoder(response).Encode(data)
}

func main() {
	database.InitDataBase()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1", Home).Methods("GET")
	config.UserApiRouting(router)

	log.Fatal(http.ListenAndServe((":" + strconv.Itoa(portNumber)), router))
}
