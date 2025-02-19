package domain

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilmsg/improved-couscous/note-zero/model"
)

type NoteHandler interface {
	CreateNote(w http.ResponseWriter, r *http.Request)
	ListNote(w http.ResponseWriter, r *http.Request)
	GetNote(w http.ResponseWriter, r *http.Request)
	UpdateNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}

type NoteService interface {
	CreateNote(note model.Note) error
	ListNote() ([]model.Note, error)
	GetNote(id uuid.UUID) (model.Note, error)
	UpdateNote(note model.Note) error
	DeleteNote(id uuid.UUID) error
}

type NoteRepo interface {
	CreateNote(note model.Note) error
	ListNote() ([]model.Note, error)
	GetNote(id uuid.UUID) (model.Note, error)
	UpdateNote(note model.Note) error
	DeleteNote(id uuid.UUID) error
}
