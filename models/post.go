package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Body          string `gorm:"not null"`
	PublishedDate time.Time
}
