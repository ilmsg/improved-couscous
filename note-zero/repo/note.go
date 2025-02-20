package repo

import (
	"github.com/google/uuid"
	"github.com/ilmsg/improved-couscous/note-zero/domain"
	"github.com/ilmsg/improved-couscous/note-zero/model"
	"gorm.io/gorm"
)

type noteRepo struct {
	db *gorm.DB
}

// CreateNote implements domain.NoteRepo.
func (n *noteRepo) CreateNote(note *model.Note) error {
	return n.db.Create(&note).Error
}

// DeleteNote implements domain.NoteRepo.
func (n *noteRepo) DeleteNote(id uuid.UUID) error {
	return n.db.Delete(&model.Note{}, id).Error
}

// GetNote implements domain.NoteRepo.
func (n *noteRepo) GetNote(id uuid.UUID) (model.Note, error) {
	var note model.Note
	err := n.db.First(&note, id).Error
	return note, err
}

// ListNote implements domain.NoteRepo.
func (n *noteRepo) ListNote() ([]model.Note, error) {
	var notes []model.Note
	err := n.db.Find(&notes).Error
	return notes, err
}

// UpdateNote implements domain.NoteRepo.
func (n *noteRepo) UpdateNote(note model.Note) error {
	return n.db.Updates(note).Error
}

func NewRepoNote(db *gorm.DB) domain.NoteRepo {
	return &noteRepo{db}
}
