package service

import (
	"github.com/google/uuid"
	"github.com/ilmsg/improved-couscous/note-zero/domain"
	"github.com/ilmsg/improved-couscous/note-zero/model"
)

type noteService struct {
	repo domain.NoteRepo
}

// CreateNote implements domain.NoteService.
func (n *noteService) CreateNote(note model.Note) error {
	panic("unimplemented")
}

// DeleteNote implements domain.NoteService.
func (n *noteService) DeleteNote(id uuid.UUID) error {
	panic("unimplemented")
}

// GetNote implements domain.NoteService.
func (n *noteService) GetNote(id uuid.UUID) (model.Note, error) {
	panic("unimplemented")
}

// ListNote implements domain.NoteService.
func (n *noteService) ListNote() ([]model.Note, error) {
	panic("unimplemented")
}

// UpdateNote implements domain.NoteService.
func (n *noteService) UpdateNote(note model.Note) error {
	panic("unimplemented")
}

func NewServiceNote(repo domain.NoteRepo) domain.NoteService {
	return &noteService{repo}
}
