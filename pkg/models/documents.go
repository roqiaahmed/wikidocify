package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
