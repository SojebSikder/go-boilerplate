package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Country         *string    `json:"country"`
	Name            *string    `json:"name"`
	ApprovedAt      *time.Time `json:"approved_at"`
	Availability    *string    `json:"availability"`
	Email           *string    `json:"email"`
	Username        *string    `json:"username"`
	PhoneNumber     *string    `json:"phone_number"`
	FirstName       *string    `json:"first_name"`
	LastName        *string    `json:"last_name"`
	Password        *string    `json:"-"`
	TwoFactorSecret *string    `json:"-"`
	Avatar          *string    `json:"avatar"`
	ZipCode         *string    `json:"zip_code"`
	State           *string    `json:"state"`
	City            *string    `json:"city"`
	Address         *string    `json:"address"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Gender          *string    `json:"gender"`
	DateOfBirth     *time.Time `json:"date_of_birth"`
	BillingID       *string    `json:"billing_id"`
	BaseModel
	Type               string `gorm:"default:user" json:"type"`
	IsTwoFactorEnabled int    `gorm:"default:0" json:"is_two_factor_enabled"`
	Status             int    `gorm:"default:1" json:"status"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
