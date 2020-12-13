package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgresql://postgres:local-password@localhost/postgres

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "host=localhost user=postgres password=local-password dbname=users port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Stevie"})

}
