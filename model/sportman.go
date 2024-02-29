package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Sportmen struct {
	UserId           uuid.UUID `gorm:"primaryKey;type:uuid;" json:"user" validate:"required"`
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
	Sport            []SportmenSport `gorm:"foreignKey:UserId" validate:"dive"`
}

type SportmenSport struct {
	ID     uint      `gorm:"primaryKey;"  json:"id"`
	UserId uuid.UUID `gorm:"type:uuid;" json:"user"`
	Sport  string    `json:"sport" validate:"oneof=basketball cycling"`
}
