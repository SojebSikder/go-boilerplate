package ORM

import (
	"fmt"

	"github.com/sojebsikder/go-boilerplate/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init initializes the ORM
func Init(ctg config.Config) *gorm.DB {
	db, err = gorm.Open(postgres.Open(ctg.Database.DatabaseURL), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
	}

	fmt.Println("Connected to the database")
	return db
}

func GetDB() *gorm.DB {
	return db
}
