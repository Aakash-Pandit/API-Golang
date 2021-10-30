package patients

import (
	"encoding/json"
	"net/http"
	"time"

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

////////////////////////////////////////////// Patient ////////////////////////////////////////////////////

func GetAllPatients(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var patients []Patient
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to Connect DataBase")
	}
	defer db.Close()

	db.Find(&patients)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(patients)
}

func GetPatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var patient Patient
	db.Find(&patient, "id = ?", params["id"])

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(patient)
}

func CreatePatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	var patient Patient

	err := json.NewDecoder(request.Body).Decode(&patient)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to Connect Database")
	}
	defer db.Close()

	patient.Created = time.Now()
	patient.Modified = time.Now()

	db.Create(&patient)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(patient)
}

func DeletePatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(request)

	var patient Patient

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to Connect Database")
	}
	defer db.Close()

	db.Find(&patient, "id = ?", params["id"])
	db.Delete(&patient)

	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(map[string]string{})
}

func UpdatePatient(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var patient Patient
	db.Find(&patient, "id = ?", params["id"])

	err = json.NewDecoder(request.Body).Decode(&patient)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	patient.Modified = time.Now()
	db.Save(&patient)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(patient)
}

////////////////////////////////////////////// Medicine ////////////////////////////////////////////////////

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

func GetMedicine(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var medicine Medicine
	db.Find(&medicine, "id = ?", params["id"])
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(medicine)
}

func CreateMedicine(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var medicine Medicine

	err := json.NewDecoder(request.Body).Decode(&medicine)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	medicine.Created = time.Now()
	medicine.Modified = time.Now()

	db.Create(&medicine)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(medicine)
}

func DeleteMedicine(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var medicine Medicine
	db.Find(&medicine, "id = ?", params["id"])
	db.Delete(&medicine)

	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(map[string]string{})
}

func UpdateMedicine(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	defer db.Close()

	var medicine Medicine
	db.Find(&medicine, "id = ?", params["id"])
	err = json.NewDecoder(request.Body).Decode(&medicine)
	if err != nil {
		json.NewEncoder(response).Encode(err)
		return
	}
	medicine.Modified = time.Now()
	db.Save(&medicine)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&medicine)
}
