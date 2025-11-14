package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StringArray es un tipo personalizado para manejar arrays de strings como JSON en la BD
type StringArray []string

// Scan implementa la interfaz sql.Scanner para leer desde la BD
func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan StringArray: value is not []byte")
	}

	if len(bytes) == 0 {
		*s = []string{}
		return nil
	}

	var arr []string
	if err := json.Unmarshal(bytes, &arr); err != nil {
		return err
	}

	*s = arr
	return nil
}

// Value implementa la interfaz driver.Valuer para escribir a la BD
func (s StringArray) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

type Vehicle struct {
	ID        uuid.UUID   `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Titulo    string      `json:"vehicle_title" gorm:"type:char(100);not null"` //titulo de publicacion
	Brand     string      `json:"vehicle_brand" gorm:"type:char(30);not null"`
	Model     string      `json:"vehicle_model" gorm:"type:char(30);not null"`
	Price     float64     `json:"vehicle_price" gorm:"type:float;not null"`
	Year      int         `json:"vehicle_year" gorm:"type:int;not null"`
	Version   string      `json:"vehicle_version" gorm:"type:char(30);not null"`
	State     string      `json:"vehicle_state" gorm:"type:char(30);not null"`
	Image     StringArray `json:"vehicle_image" gorm:"type:json"`
	CreatedAt *time.Time  `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
}

func (v *Vehicle) BeforeCreate(tx *gorm.DB) (err error) {

	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}
