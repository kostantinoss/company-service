package main

import (
	"fmt"
	"konstantinos/company-service/internal/database"
	"konstantinos/company-service/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting server")

	db, err := database.SetUpDB()
	if err != nil {
		fmt.Println("database not setup")
	}

	_, err = database.MigrateDB(db)
	if err != nil {
		fmt.Println("Migration failed")
	}

	router := mux.NewRouter()
	
	router.HandleFunc("/create", handlers.Create(db)).Methods("POST")
	router.HandleFunc("/delete/{id}", handlers.Delete(db)).Methods("DELETE")
	router.HandleFunc("/update/{id}", handlers.Patch(db)).Methods("PUT")
	router.HandleFunc("/get/{id}", handlers.Get(db)).Methods("GET")
	
	http.ListenAndServe(":8880", router)
}