package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBodyForUpdateData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Email string `json:"email"`
}

func updateDataDiri(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start Update Data Diri")
	var requestBodyForUpdate RequestBodyForUpdateData
	var count int
	errParse := json.NewDecoder(r.Body).Decode(&requestBodyForUpdate)
	if errParse != nil {
		http.Error(w, errParse.Error(), http.StatusBadRequest)
		return
	}

	rows := db.QueryRow("SELECT COUNT(*) FROM DATADIRI WHERE id = ?", requestBodyForUpdate.ID)
	err := rows.Scan(&count)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count <= 0 {
		errorResponse := ErrorResponse{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: "ID Not Found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	_, errUpdate := db.Exec("UPDATE DATADIRI SET name = ?, age = ?, email = ? WHERE id =?", requestBodyForUpdate.Name, requestBodyForUpdate.Age, requestBodyForUpdate.Email, requestBodyForUpdate.ID)
	if errUpdate != nil {
		log.Fatal(err)
		errorResponse := ErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "FAILED UPDATE DATADIRI",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	errorResponse := ErrorResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "Succes Update Data Diri",
	}
	json.NewEncoder(w).Encode(errorResponse)

}
