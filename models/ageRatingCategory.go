package models

import (
	"time"
)

type (
	// AgeRatingCategory
	AgeRatingCategory struct {
		ID          int       `gorm:"primary_key" json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Movies      []Movie   `json:"-"` // untuk menandakan relasi dengan Movie
	}
)
