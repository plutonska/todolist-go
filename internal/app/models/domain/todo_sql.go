package domain

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	UUID      string `gorm:"type:char(36);not null;unique"`
	ID        uint   `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Item      string         `gorm:"text;not null"`
	Done      bool           `gorm:"bool;default:false"`
}
