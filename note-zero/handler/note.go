package handler

import (
	"net/http"

	"github.com/ilmsg/improved-couscous/note-zero/domain"
)

type noteHandler struct {
	srv domain.NoteService
}

// CreateNote implements domain.NoteHandler.
func (n *noteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// DeleteNote implements domain.NoteHandler.
func (n *noteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetNote implements domain.NoteHandler.
func (n *noteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// ListNote implements domain.NoteHandler.
func (n *noteHandler) ListNote(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// UpdateNote implements domain.NoteHandler.
func (n *noteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewHandlerNote(srv domain.NoteService) domain.NoteHandler {
	return &noteHandler{srv}
}
