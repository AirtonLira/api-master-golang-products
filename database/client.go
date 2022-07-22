package database

import (
	"api-crud/entities"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect(connectionString string) {
	inst, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")

	Instance = inst
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database Migration Completed...")
}
