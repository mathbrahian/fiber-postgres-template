package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Sportmen struct {
	ID               uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4();" json:"id"`
	User             uuid.UUID `json:"user" validate:"required"`
	Name             string    `json:"name" validate:"required"`
	LastName         string    `json:"last_name" validate:"required"`
	Age              int       `json:"age" validate:"required"`
	Weight           int       `json:"weight" validate:"required"`
	Height           int       `json:"height" validate:"required"`
	CountryBirth     string    `json:"country_birth" validate:"required"`
	CityBirth        string    `json:"city_birth" validate:"required"`
	CountryResidence string    `json:"country_residence" validate:"required"`
	CityResidence    string    `json:"city_residence" validate:"required"`
	LengthResidence  int       `json:"length_residence" validate:"required"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Sport            []SportmenSport `gorm:"foreignKey:User" validate:"dive"`
}

type SportmenSport struct {
	ID    int       `gorm:"primary_key" json:"id"`
	User  uuid.UUID `gorm:"type:uuid;" json:"user"`
	Sport string    `json:"sport" validate:"oneof=basketball cycling"`
}
