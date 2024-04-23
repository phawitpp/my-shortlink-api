package model

import (
	"time"

	"github.com/google/uuid"
)

type LinkModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	URL       string    `gorm:"size:255;not null" json:"url"`
	ShortLink string    `gorm:"size:255;not null;unique" json:"short_link"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
