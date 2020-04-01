package model

import "time"

// Model represents meta data of entity.
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id" structs:"id"`
	CreatedAt time.Time  `json:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" structs:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt" structs:"deletedAt"`
}
