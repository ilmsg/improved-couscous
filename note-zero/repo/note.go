package repo

import (
	"github.com/ilmsg/improved-couscous/note-zero/domain"
	"gorm.io/gorm"
)

type noteRepo struct {
	db *gorm.DB
}

func NewRepoNote(db *gorm.DB) domain.NoteRepo {
	return &noteRepo{db}
}
