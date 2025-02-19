package model

import "github.com/google/uuid"

type Note struct {
	ID   uuid.UUID `json:"id"`
	Note string    `json:"note"`
}
