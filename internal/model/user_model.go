package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Model

	Status       int        `gorm:"default:1" json:"status"`
	ApprovedAt   *time.Time `json:"approved_at"`
	Availability *string    `json:"availability"`

	Email     *string `json:"email"`
	Username  *string `json:"username"`
	Name      *string `json:"name"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`

	Password *string `json:"-"`

	Avatar             *string    `json:"avatar"`
	PhoneNumber        *string    `json:"phone_number"`
	Country            *string    `json:"country"`
	State              *string    `json:"state"`
	City               *string    `json:"city"`
	Address            *string    `json:"address"`
	ZipCode            *string    `json:"zip_code"`
	Gender             *string    `json:"gender"`
	DateOfBirth        *time.Time `json:"date_of_birth"`
	BillingID          *string    `json:"billing_id"`
	Type               string     `gorm:"default:user" json:"type"`
	EmailVerifiedAt    *time.Time `json:"email_verified_at"`
	IsTwoFactorEnabled int        `gorm:"default:0" json:"is_two_factor_enabled"`
	TwoFactorSecret    *string    `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
