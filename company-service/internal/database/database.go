package database

import (
	"fmt"
	"konstantinos/company-service/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func SetUpDB() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("companies.db"), &gorm.Config{})
  	if err != nil {
    	fmt.Printf("failed to connect database, with err: %v", err.Error())
		return nil, err
 	}

	return db, nil
}

func MigrateDB(db *gorm.DB) (ok bool, err error) {
	err = db.AutoMigrate(&model.Company{})
	if err != nil {
		fmt.Printf("failed to migrate model, with err: %v", err.Error())
		return false, err
	}

	return ok, err
}