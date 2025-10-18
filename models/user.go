package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Name      string     `json:"user_name" gorm:"type:char(60);not null"`
	Email     string     `json:"user_email" gorm:"type:char(120);not null"`
	Phone     string     `json:"user_phone" gorm:"type:char(10);not null"`
	Password  string     `json:"user_password" gorm:"type: char (250); not null"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
