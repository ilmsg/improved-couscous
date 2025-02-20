package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	ID   uuid.UUID `json:"id"`
	Note string    `json:"note"`
}

func (n *Note) BeforeCreate(db *gorm.DB) error {
	if n.ID.String() == "00000000-0000-0000-0000-000000000000" {
		n.ID = uuid.New()
	}
	return nil
}
