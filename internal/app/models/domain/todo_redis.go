package domain

import "time"

type Todos struct {
	UUID      string     `json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Item      string     `json:"item"`
	Done      bool       `json:"done"`
}
