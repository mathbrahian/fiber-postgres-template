package model

import "time"

import uuid "github.com/google/uuid"

type Note struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4();" json:"id"`
	User      uuid.UUID `json:"user"`
	Note      string    `json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
