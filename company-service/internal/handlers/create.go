package handlers

import (
	"encoding/json"
	"fmt"
	"konstantinos/company-service/internal/model"
	"net/http"

	"gorm.io/gorm"
)

func Create(db *gorm.DB) http.HandlerFunc {
	fmt.Println("create called")
	handler := func (w http.ResponseWriter, req *http.Request) {
		var company model.Company
		dec := json.NewDecoder(req.Body)
    	dec.DisallowUnknownFields()
		err := dec.Decode(&company)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
        	return
		}

		db.Create(&company)
		w.WriteHeader(http.StatusCreated)
    	json.NewEncoder(w).Encode(company)
	}

	return handler
}
