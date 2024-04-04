package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBody struct {
	ID int `json:"id"`
}

func getDataDiriById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start Get Data Diri By ID")
	var requestBody RequestBody
	errParse := json.NewDecoder(r.Body).Decode(&requestBody)
	if errParse != nil {
		http.Error(w, errParse.Error(), http.StatusBadRequest)
		return
	}

	var dataDiri DataDiri
	rows := db.QueryRow("SELECT * FROM DATADIRI WHERE id = ?", requestBody.ID)
	err := rows.Scan(&dataDiri.ID, &dataDiri.Name, &dataDiri.Age, &dataDiri.Email, &dataDiri.CreatedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			errorResponse := ErrorResponse{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: "ID Not Found",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(dataDiri)
}
