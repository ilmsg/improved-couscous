package note

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Note struct {
	ID   uuid.UUID `json:"id"`
	Note string    `json:"note"`
}

func (n *Note) Hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func (n *Note) List(w http.ResponseWriter, r *http.Request) {
	notes := []Note{
		{ID: uuid.New(), Note: "Take Note 1"},
		{ID: uuid.New(), Note: "Take Note 2"},
	}
	json.NewEncoder(w).Encode(notes)
}
