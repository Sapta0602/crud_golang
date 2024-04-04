package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBodyCreateData struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Email string `json:"email"`
}

func createDataDiri(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start Create Data Diri")
	var requestBodyCreateData RequestBodyCreateData
	errParse := json.NewDecoder(r.Body).Decode(&requestBodyCreateData)
	if errParse != nil {
		http.Error(w, errParse.Error(), http.StatusBadRequest)
		return
	}
	_, err := db.Exec("INSERT INTO DATADIRI (name, age, email) VALUES (?,?,?)", requestBodyCreateData.Name, requestBodyCreateData.Age, requestBodyCreateData.Email)
	if err != nil {
		log.Fatal(err)
	}
	errorResponse := ErrorResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "Success Created Data",
	}

	json.NewEncoder(w).Encode(errorResponse)

}
