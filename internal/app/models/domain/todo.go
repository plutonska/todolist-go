package domain

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	UUID      string         `gorm:"type:char(36);primaryKey;not null;unique"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Item      string         `gorm:"text;not null"`
	Done      bool           `gorm:"bool;default:false"`
}
