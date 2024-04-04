package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAllDataDiriHandler(w http.ResponseWriter, r *http.Request) {
	var dataDiris []DataDiri
	rows, err := db.Query("SELECT * FROM DATADIRI")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dataDiri DataDiri
		err := rows.Scan(&dataDiri.ID, &dataDiri.Name, &dataDiri.Age, &dataDiri.Email, &dataDiri.CreatedDate)
		if nil != err {
			log.Fatal(err)
		}
		dataDiris = append(dataDiris, dataDiri)
	}

	json.NewEncoder(w).Encode(dataDiris)

}
