package model

import (
	"time"

	"github.com/sojebsikder/go-boilerplate/pkg/ORM"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ORM.Model

	Name     string `gorm:"nullable" json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Type     string `gorm:"default:user;nullable" json:"type"`

	Status int `gorm:"default:1" json:"status"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now().UTC()
	return
}
