package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Titulo    string     `json:"vehicle_title" gorm:"type:char(100);not null"` //titulo de publicacion
	Brand     string     `json:"vehicle_brand" gorm:"type:char(30);not null"`
	Model     string     `json:"vehicle_model" gorm:"type:char(30);not null"`
	Price     float64    `json:"vehicle_price" gorm:"type:float;not null"`
	Year      int        `json:"vehicle_year" gorm:"type:int;not null"`
	Version   string     `json:"vehicle_version" gorm:"type:char(30);not null"`
	State     string     `json:"vehicle_state" gorm:"type:char(30);not null"`
	Image     []string   `json:"vehicle_image" gorm:"type:json"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (v *Vehicle) BeforeCreate(tx *gorm.DB) (err error) {

	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}
