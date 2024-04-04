package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBodyForDelete struct {
	ID int `json:"id"`
}

func deleteDataDiriByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start Delete Data Diri By ID")
	var requestBodyForDelete RequestBodyForDelete
	var count int
	errParse := json.NewDecoder(r.Body).Decode(&requestBodyForDelete)
	if errParse != nil {
		http.Error(w, errParse.Error(), http.StatusBadRequest)
		return
	}

	rows := db.QueryRow("SELECT COUNT(*) FROM DATADIRI WHERE id = ?", requestBodyForDelete.ID)
	err := rows.Scan(&count)
	if err != nil {
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
	_, errDelete := db.Exec("DELETE FROM DATADIRI WHERE id = ?", requestBodyForDelete.ID)
	if errDelete != nil {
		log.Fatal(err)
		errorResponse := ErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "FAILED DELETE DATADIRI",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	errorResponse := ErrorResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "Succes Delete Data Diri",
	}
	json.NewEncoder(w).Encode(errorResponse)
}
