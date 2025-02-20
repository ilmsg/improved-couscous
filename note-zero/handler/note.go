package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ilmsg/improved-couscous/note-zero/domain"
	"github.com/ilmsg/improved-couscous/note-zero/model"
)

type noteHandler struct {
	srv domain.NoteService
}

// CreateNote implements domain.NoteHandler.
func (n *noteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	if err := n.srv.CreateNote(&note); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(note)
}

// DeleteNote implements domain.NoteHandler.
func (n *noteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	id, _ := uuid.Parse(paramId)

	if err := n.srv.DeleteNote(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(domain.MessageResponse{Message: fmt.Sprintf("delete %s successfully", id)})
}

// GetNote implements domain.NoteHandler.
func (n *noteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	id, _ := uuid.Parse(paramId)

	note, err := n.srv.GetNote(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(note)
}

// ListNote implements domain.NoteHandler.
func (n *noteHandler) ListNote(w http.ResponseWriter, r *http.Request) {
	notes, err := n.srv.ListNote()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(notes)
}

// UpdateNote implements domain.NoteHandler.
func (n *noteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	id, _ := uuid.Parse(paramId)

	var noteReq struct {
		Note string `json:"note"`
	}
	if err := json.NewDecoder(r.Body).Decode(&noteReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	updateNote, err := n.srv.GetNote(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	updateNote.Note = noteReq.Note
	if err := n.srv.UpdateNote(updateNote); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(updateNote)
}

func NewHandlerNote(srv domain.NoteService) domain.NoteHandler {
	return &noteHandler{srv}
}
