package models

import "time"

type Author struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null" validate:"required" json:"name"`
	Birthdate time.Time `gorm:"not null" validate:"required" time_format:"2006-01-02"`
	
}
