package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	if err != nil {
		panic("Database Failed")
	}

	DB = database
}
