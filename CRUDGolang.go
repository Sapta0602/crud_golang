package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type DataDiri struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	Email       string `json:"email"`
	CreatedDate string `json:"created_at"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

var db *sql.DB

func main() {
	fmt.Println("Test Create CRUD")

	var err error
	db, err = sql.Open("mysql", "root:P@ssw0rd@tcp(localhost:3306)/MYDATABASE")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connect To Database")

	r := mux.NewRouter()
	r.HandleFunc("/datas", getAllDataDiriHandler).Methods("GET")
	r.HandleFunc("/dataByID", getDataDiriById).Methods("GET")
	r.HandleFunc("/createDataDiri", createDataDiri).Methods("POST")
	r.HandleFunc("/deleteDataDiriByID", deleteDataDiriByID).Methods("DELETE")
	r.HandleFunc("/updateDataDiri", updateDataDiri).Methods("PUT")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
