package handlers

import (
	"encoding/json"
	"konstantinos/company-service/internal/model"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Patch(db *gorm.DB) http.HandlerFunc {
	handler := func (w http.ResponseWriter, req *http.Request) {
		var company model.Company
		err := json.NewDecoder(req.Body).Decode(&company)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vars := mux.Vars(req)
		companyID := vars["id"]

		db.Update(companyID, &company)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(company)
	}

	return handler
}