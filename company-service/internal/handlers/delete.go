package handlers

import (
	"fmt"
	"konstantinos/company-service/internal/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) http.HandlerFunc {
	fmt.Println("delete called")
	handler := func (w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		companyID := vars["id"]
		companyID_uuid, _ := uuid.FromBytes([]byte(companyID))

		db.Delete(&model.Company{ID: companyID_uuid})

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Company deleted successfully"}`))
	}

	return handler
}