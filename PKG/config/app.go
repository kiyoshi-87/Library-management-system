package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "ENTER YOUR CONNECTION STRING HERE"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Checking the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Log success message after successfully obtaining the connection
	log.Println("Successfully Connected to the MySQL db!")

	// Defer the Close() call to ensure the connection is closed when the surrounding function returns
	defer sqlDB.Close()

	return db
}
