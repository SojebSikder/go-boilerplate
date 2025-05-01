package orm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init initializes the ORM
func Init(configString string) {
	db, err = gorm.Open(postgres.Open(configString), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	fmt.Println("Connected to the database")
}

func GetDB() *gorm.DB {
	return db
}
