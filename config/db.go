package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(){
	var err error
	DB, err = gorm.Open(sqlite.Open("/app/app.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed : %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed : %v", err)
	}

	_, err = sqlDB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatalf("Failed : %v", err)
	}

	log.Println("DB connected and fk")
}