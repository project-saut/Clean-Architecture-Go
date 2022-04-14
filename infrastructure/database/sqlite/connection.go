package database

import (
	"Belajar-Go-Echo/domains/users"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("DB Error to Connect")
		return nil
	}

	e := DB.AutoMigrate(
		users.User{},
	)

	if e != nil {
		return nil
	}

	return DB
}
