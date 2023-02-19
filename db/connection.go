package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DSN = "host=localhost user=fazt password=mysecretpassword dbname=gorm port=5432"
	DB  *gorm.DB
)

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successfully")
	
}
