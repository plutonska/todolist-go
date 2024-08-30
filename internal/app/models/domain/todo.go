package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Item string `gorm:"text;not null"`
	Done bool   `gorm:"bool;default:false"`
}
