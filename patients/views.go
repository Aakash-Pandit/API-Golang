package patients

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var patients []Patient

func CreateDummyPatients() {
	patients = append(patients, Patient{ID: uuid.New(), FirstName: "Aakash", LastName: "Pandit", Email: "aakashpandit366@gmail.com", Contact: "8698410175"})
	patients = append(patients, Patient{ID: uuid.New(), FirstName: "Siddhesh", LastName: "Pandit", Email: "Sid@gmail.com", Contact: "1234543210"})
	patients = append(patients, Patient{ID: uuid.New(), FirstName: "Rasika", LastName: "Pandit", Email: "ras@gmail.com", Contact: "1234543210"})
	return
}

func GetAllPatients(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader((200))
	json.NewEncoder(response).Encode(patients)
}

func GetPatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)
	for _, item := range patients {
		if item.ID.String() == params["id"] {
			response.WriteHeader((200))
			json.NewEncoder(response).Encode(item)
			return
		}
	}
	response.WriteHeader((200))
	json.NewEncoder(response).Encode(map[string]string{})
}

func CreatePatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	var patient Patient
	err := json.NewDecoder(request.Body).Decode(&patient)

	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}
	patient.BeforeCreate()
	patients = append(patients, patient)
	response.WriteHeader((201))
	json.NewEncoder(response).Encode(patient)
}

func DeletePatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)

	for index, item := range patients {
		if item.ID.String() == params["id"] {
			patients = append(patients[:index], patients[index+1:]...)
			response.WriteHeader((204))
			json.NewEncoder(response).Encode(map[string]string{})
			return
		}
	}
	response.WriteHeader((400))
	json.NewEncoder(response).Encode(map[string]string{"details": "Bad Request"})
}

func UpdatePatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range patients {

		var patient Patient
		if item.ID.String() == params["id"] {
			patients = append(patients[:index], patients[index+1:]...)
			err := json.NewDecoder(request.Body).Decode(&patient)
			if err != nil {
				response.WriteHeader((400))
				json.NewEncoder(response).Encode(err)
				return
			}
			response.WriteHeader((200))
			id, _ := uuid.Parse(params["id"])
			patient.ID = id
			patients = append(patients, patient)
			json.NewEncoder(response).Encode(patient)
			return
		}
	}

	response.WriteHeader((400))
	json.NewEncoder(response).Encode(map[string]string{})
}

func (patient *Patient) BeforeCreate() Patient {
	(*patient).ID = uuid.New()
	return *patient
}

func GetAllMedicines(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var medicines []Medicine

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Find(&medicines)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(medicines)
}

func CreateMedicine(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var medicine = Medicine{Name: "Medicine A", Cost: 10}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(&medicine)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(medicine)
}
