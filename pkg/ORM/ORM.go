package ORM

import (
	"errors"
	"fmt"

	"github.com/sojebsikder/go-boilerplate/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init initializes the ORM
func Init(ctg *config.Config) (*gorm.DB, error) {
	db, err = gorm.Open(postgres.Open(ctg.Database.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Failed to connect to the database:" + err.Error())
	}

	fmt.Println("Connected to the database")
	return db, nil
}
