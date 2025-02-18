package note

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList_Hi(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	note := Note{}
	note.Hi(w, r)

	exp := "hi"
	if w.Body.String() != exp {
		t.Errorf("expected '%s', got '%s'", exp, w.Body.String())
	}
}

func TestList(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	note := Note{}
	note.List(w, r)

	var notes []Note
	err := json.Unmarshal(w.Body.Bytes(), &notes)
	if err != nil {
		t.Errorf("expected json.Unmarshal error to be nil")
	}

	expect := 2
	if len(notes) != 2 {
		t.Errorf("expected notes '%d', got '%d'", expect, len(notes))
	}
}
